INSERT INTO users (
    id,
    email,
    password,
    password_salt,
    role,
    is_active,
    requested_role,
    verification_status,
    created_at,
    updated_at,
    is_deleted
) VALUES (
             '8aa7d90e-d32f-4a09-ac66-c1209311fd04',
             'hanif@kpupusat.go.id',         -- adjust email as appropriate
             'e6d0e46a33a438c516dcaf4d27cadccfed1b68282a855ace4f16ce148fa738ed7a20fd4ba8ba38c9decb47e41a485937dc0b8e930385ff4479582cbb077a52bf',         -- replace with actual hashed password
             'y%&NbSn',                       -- replace with actual salt
             'kpu_pusat',
             true,
             NULL,
             'approved',
             '2025-05-22 12:19:09.092279+07',
             '2025-05-22 12:19:09.092279+07',
             false
         );

INSERT INTO kpu_provinsi (
    id,
    user_id,
    name,
    username,
    address,
    region,
    is_active,
    photo_path,
    telephone,
    registered_at,
    created_at,
    updated_at,
    is_deleted
) VALUES (
             '6e7545e4-12fb-42a2-b206-4b921f38700e',
          '8aa7d90e-d32f-4a09-ac66-c1209311fd04',
             'KPU Pusat',
             'Hanif',
             '0x6970322e22A7C880A9385E80017ff7A9489f683b',
             'Indonesia',
             true,
             '',
             '+6281234567890',
             '2025-05-22 12:19:09.092279+07',
             '2025-05-22 12:19:09.092279+07',
             '2025-05-22 12:19:09.092279+07',
                false
);