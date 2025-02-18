// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Votechain {

    struct KPUBranch{
        string name;
        address branchAddress;
        bool isActive;
        string region;
    }

    struct Voter{
        string ktp;
        address voterAddress;
        bool hasVoted;
        string region;
        bool isRegistered;
    }

    struct Candidate {
        uint256 id;
        string name;
        uint256 voteCount;
        bool isActive;
    }

    address public kpuAdmin;
    mapping(address => KPUBranch) public kpuBranches;
    mapping(string => Voter) public voters;
    mapping(uint256 => Candidate) public candidates;
    uint256 public candidateCount;
    bool public votingActive;

    KPUBranch[]public kpuBranchAddresses;
    Voter[] public voterAddresses;

    event KPUBranchRegistered(address branchAddress, string name, string region);
    event VoterRegistered(string ktp, address voterAddress, string region);
    event VoteCasted(string ktp, uint256 candidateId);
    event CandidateAdded(uint256 candidateId, string name);
    event VotingStatusChange(bool isActive);

    modifier onlyKpuAdmin(){
        require(msg.sender == kpuAdmin, "Only KPU admin can do this action");
        _;
    }

    modifier onlyKpuBranch(){
        require(kpuBranches[msg.sender].isActive, "Only active KPU branch can do this action");
        _;
    }

    modifier votingIsActive(){
        require(votingActive, "Voting session is not active");
        _;
    }
    constructor(){
        kpuAdmin = msg.sender;
        votingActive = false;
    }

    function getAllKPUBranches() public view returns (KPUBranch[] memory) {
        return kpuBranchAddresses;
    }

    function getKPUBranchAddress(address branchAddress) public view returns (KPUBranch memory) {
        KPUBranch memory branch = kpuBranches[branchAddress];
        return KPUBranch(branch.name, branch.branchAddress, branch.isActive, branch.region);
    }

    function getAllVoter() public view returns (Voter[] memory) {
        return voterAddresses;
    }

    function getVoterByKTP(string memory ktp) public view returns (Voter memory) {
        return voters[ktp];
    }

    function getVoterByAddress(address voterAddress) public view returns (Voter memory) {
        for (uint i = 0; i < voterAddresses.length; i++) {
            if (voterAddresses[i].voterAddress == voterAddress) {
                return voterAddresses[i];
            }
        }
        revert("Voter not found");
    }

    function getVotersByRegion(string memory region) public view returns (Voter[] memory) {
        uint256 voterCount = 0;

        for (uint i = 0; i < voterAddresses.length; i++) {
            if (keccak256(abi.encodePacked(voterAddresses[i].region)) == keccak256(abi.encodePacked(region))) {
                voterCount++;
            }
        }

        Voter[] memory matchingVoters = new Voter[](voterCount);
        uint256 index = 0;

        for (uint i = 0; i < voterAddresses.length; i++) {
            if (keccak256(abi.encodePacked(voterAddresses[i].region)) == keccak256(abi.encodePacked(region))) {
                matchingVoters[index] = voterAddresses[i];
                index++;
            }
        }

        return matchingVoters;
    }

    function addKpuBranch (address branchAddress, KPUBranch memory kpuInstance) public onlyKpuAdmin{
        require(!kpuBranches[branchAddress].isActive, "Branch already registered!!!");

        kpuBranches[branchAddress] = kpuInstance;
    }

    function registerKPUBranch(address branchAddress, string memory name, string memory region) public onlyKpuAdmin  {
        require(!kpuBranches[branchAddress].isActive, "Branch already registered!!!");

        KPUBranch memory newKPUBranch = (
            KPUBranch({
            name : name,
            branchAddress: branchAddress,
            isActive: true,
            region: region
        })
        );

        kpuBranches[branchAddress] = newKPUBranch;
        kpuBranchAddresses.push(newKPUBranch);

        emit KPUBranchRegistered(branchAddress, name,region);
    }


    function deactivateKPUBranch(address branchAddress) public onlyKpuAdmin{
        require(kpuBranches[branchAddress].isActive, "Branch is not active");
        kpuBranches[branchAddress].isActive = false;
    }

    function addCandidate(string memory name) public onlyKpuAdmin  {
        candidateCount++;
        candidates[candidateCount] = Candidate({
            id: candidateCount,
            name: name,
            voteCount: 0,
            isActive : true
        });

        emit CandidateAdded(candidateCount, name);
    }

    function setVotingStatus(bool status) public onlyKpuAdmin  {
        votingActive = status;
        emit VotingStatusChange(status);
    }

    function registerVoter(
        string memory ktp,
        address voterAddress
    ) public onlyKpuBranch {
        require(!voters[ktp].isRegistered, "Voter already registered!");

        Voter memory voter = (Voter({
            ktp : ktp,
            voterAddress : voterAddress,
            hasVoted: false,
            region: kpuBranches[msg.sender].region,
            isRegistered: true
        }));

        voters[ktp] = Voter({
            ktp : ktp,
            voterAddress : voterAddress,
            hasVoted: false,
            region: kpuBranches[msg.sender].region,
            isRegistered: true
        });

        voterAddresses.push(voter);

        emit VoterRegistered(ktp, voterAddress, kpuBranches[msg.sender].region);
    }

    function vote(string memory ktp, uint256 candidateId) public votingIsActive{
        Voter storage voter = voters[ktp];
        require(voter.isRegistered, "Voter not registered!!");
        require(!voter.hasVoted, "already voted");
        require(voter.voterAddress == msg.sender, "not authorized to vote with this KTP");
        require(candidates[candidateId].isActive, "Invalid Candidate");

        voter.hasVoted = true;
        candidates[candidateId].voteCount++;

        emit VoteCasted(ktp, candidateId);

    }

    function getCandidateCount() public view returns (uint256)  {
        return candidateCount;
    }

    function getVoterStatus(string memory ktp) public view returns (bool isRegistered, bool hasVoted, string memory region){
        Voter memory voter = voters[ktp];
        return (voter.isRegistered, voter.hasVoted, voter.region);
    }

    function getCandidatesVotes(uint256 candidateId) public view returns (uint256)  {
        require(candidateId > 0 && candidateId <= candidateCount, "Invalid Candidate ID");
        return candidates[candidateId].voteCount;
    }
    function setKpuAdmin(address newAdmin) public onlyKpuAdmin{
        kpuAdmin = newAdmin;
    }
    function getKpuAdmin() public view returns (address){
        return kpuAdmin;
    }
}
