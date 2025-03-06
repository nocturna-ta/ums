// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

contract Votechain {
    error OnlyKpuAdmin();
    error OnlyKpuBranch();
    error VotingNotActive();
    error BranchAlreadyRegistered();
    error BranchNotActive();
    error VoterAlreadyRegistered();
    error AddressAlreadyRegistered();
    error VoterNotRegistered();
    error AlreadyVoted();
    error InvalidCandidate();
    error UnauthorizedVoter();

    struct KPUBranch {
        string name;
        address branchAddress;
        bool isActive;
        string region;
    }

    struct Voter {
        string nik;
        address voterAddress;
        bool hasVoted;
        string region;
        bool isRegistered;
    }

    struct Candidate {
        string name;
        uint256 voteCount;
        bool isActive;
    }

    address public kpuAdmin;
    bool public votingActive;
    uint256 public candidateCount;

    mapping(address => KPUBranch) public kpuBranches;
    mapping(string => Voter) public voters;
    mapping(uint256 => Candidate) public candidates;
    mapping(address => string) public voterNIKByAddress;

    KPUBranch[] public kpuBranchAddresses;
    Voter[] public voterAddresses;

    event KPUBranchRegistered(address indexed branchAddress, string name, string region);
    event KPUBranchDeactivated(address indexed branchAddress);
    event VoterRegistered(string indexed nik, address indexed voterAddress, string region);
    event VoteCasted(string indexed nik, uint256 indexed candidateId);
    event CandidateAdded(uint256 indexed candidateId, string name);
    event CandidateStatusChanged(uint256 indexed candidateId, bool isActive);
    event VotingStatusChanged(bool isActive);

    modifier onlyKpuAdmin() {
        if (msg.sender != kpuAdmin) revert OnlyKpuAdmin();
        _;
    }

    modifier onlyKpuBranch() {
        if (!kpuBranches[msg.sender].isActive) revert OnlyKpuBranch();
        _;
    }

    modifier votingIsActive() {
        if (!votingActive) revert VotingNotActive();
        _;
    }

    constructor() {
        kpuAdmin = msg.sender;
    }

    function registerKPUBranch(
        address branchAddress,
        string calldata name,
        string calldata region
    ) external onlyKpuAdmin {
        if (kpuBranches[branchAddress].isActive) revert BranchAlreadyRegistered();

        kpuBranches[branchAddress] = KPUBranch(
            name,
            branchAddress,
            true,
            region
        );

        KPUBranch memory newBranch = KPUBranch(name, branchAddress, true, region);
        kpuBranches[branchAddress] = newBranch;
        kpuBranchAddresses.push(newBranch);

        emit KPUBranchRegistered(branchAddress, name, region);
    }

    function deactivateKPUBranch(address branchAddress) external onlyKpuAdmin {
        if (!kpuBranches[branchAddress].isActive) revert BranchNotActive();
        kpuBranches[branchAddress].isActive = false;
        emit KPUBranchDeactivated(branchAddress);
    }

    function registerVoter(
        string calldata nik,
        address voterAddress
    ) external onlyKpuBranch {
        if (voters[nik].isRegistered) revert VoterAlreadyRegistered();
        if (bytes(voterNIKByAddress[voterAddress]).length != 0) revert AddressAlreadyRegistered();
        Voter memory newVoter = Voter({
            nik: nik,
            voterAddress: voterAddress,
            hasVoted: false,
            region: kpuBranches[msg.sender].region,
            isRegistered: true
        });

        voters[nik] = newVoter;
        voterNIKByAddress[voterAddress] = nik;
        voterAddresses.push(newVoter);

        emit VoterRegistered(nik, voterAddress, kpuBranches[msg.sender].region);
    }

    function vote(uint256 candidateId) external votingIsActive {
        string memory nik = voterNIKByAddress[msg.sender];
        if (bytes(nik).length == 0) revert VoterNotRegistered();

        Voter storage voter = voters[nik];
        if (voter.hasVoted) revert AlreadyVoted();
        if (candidateId == 0 || candidateId > candidateCount || !candidates[candidateId].isActive)
            revert InvalidCandidate();

        voter.hasVoted = true;
        candidates[candidateId].voteCount++;

        emit VoteCasted(nik, candidateId);
    }

    function addCandidate(string calldata name) external onlyKpuAdmin {
        candidateCount++;
        candidates[candidateCount] = Candidate(name, 0, true);
        emit CandidateAdded(candidateCount, name);
    }

    function toggleCandidateActive(uint256 candidateId) external onlyKpuAdmin {
        if (candidateId == 0 || candidateId > candidateCount) revert InvalidCandidate();
        candidates[candidateId].isActive = !candidates[candidateId].isActive;
        emit CandidateStatusChanged(candidateId, candidates[candidateId].isActive);
    }

    function setVotingStatus(bool status) external onlyKpuAdmin {
        votingActive = status;
        emit VotingStatusChanged(status);
    }

    function setKpuAdmin(address newAdmin) external onlyKpuAdmin {
        kpuAdmin = newAdmin;
    }

    function getAllKPUBranches() external view returns (KPUBranch[] memory) {
        return kpuBranchAddresses;
    }

    function getBranchByAddress(address branchAddress) external view returns (KPUBranch memory) {
        if (!kpuBranches[branchAddress].isActive) revert BranchNotActive();
        return kpuBranches[branchAddress];
    }

    function getAllVoter() external view returns (Voter[] memory) {
        return voterAddresses;
    }

    function getVoterByNIK(string calldata nik) external view returns (Voter memory) {
        return voters[nik];
    }

    function getVoterByAddress(address voterAddress) external view returns (Voter memory) {
        string memory nik = voterNIKByAddress[voterAddress];
        if (bytes(nik).length == 0) revert VoterNotRegistered();
        return voters[nik];
    }

    function getVoterByRegion(string calldata region) external view returns (Voter[] memory) {
        uint256 count;
        for (uint256 i = 0; i < voterAddresses.length; i++) {
            if (keccak256(bytes(voterAddresses[i].region)) == keccak256(bytes(region))) {
                count++;
            }
        }

        Voter[] memory regionVoters = new Voter[](count);
        uint256 index;
        for (uint256 i = 0; i < voterAddresses.length; i++) {
            if (keccak256(bytes(voterAddresses[i].region)) == keccak256(bytes(region))) {
                regionVoters[index] = voterAddresses[i];
                index++;
            }
        }
        return regionVoters;
    }

    function getCandidate(uint256 candidateId) external view returns (Candidate memory) {
        if (candidateId == 0 || candidateId > candidateCount) revert InvalidCandidate();
        return candidates[candidateId];
    }

    function getAllCandidates() external view returns (Candidate[] memory) {
        Candidate[] memory allCandidates = new Candidate[](candidateCount);
        for (uint256 i = 1; i <= candidateCount; i++) {
            allCandidates[i-1] = candidates[i];
        }
        return allCandidates;
    }
}