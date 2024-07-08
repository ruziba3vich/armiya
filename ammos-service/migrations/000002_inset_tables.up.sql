BEGIN;

-- Mock data for ammos table
INSERT INTO ammos (id, name, caliber, description, type, quantity, last_update) VALUES
  (uuid_generate_v4(), 'Ammo A', '9mm', 'Standard 9mm ammo', 'Handgun', 1000, NOW()),
  (uuid_generate_v4(), 'Ammo B', '5.56mm', 'Standard 5.56mm ammo', 'Rifle', 500, NOW()),
  (uuid_generate_v4(), 'Ammo C', '7.62mm', 'Standard 7.62mm ammo', 'Rifle', 300, NOW()),
  (uuid_generate_v4(), 'Ammo D', '12 Gauge', 'Standard 12 Gauge ammo', 'Shotgun', 200, NOW()),
  (uuid_generate_v4(), 'Ammo E', '45 ACP', 'Standard 45 ACP ammo', 'Handgun', 800, NOW()),
  (uuid_generate_v4(), 'Ammo F', '40 S&W', 'Standard 40 S&W ammo', 'Handgun', 600, NOW()),
  (uuid_generate_v4(), 'Ammo G', '380 ACP', 'Standard 380 ACP ammo', 'Handgun', 700, NOW()),
  (uuid_generate_v4(), 'Ammo H', '10mm', 'Standard 10mm ammo', 'Handgun', 400, NOW()),
  (uuid_generate_v4(), 'Ammo I', '308 Win', 'Standard 308 Win ammo', 'Rifle', 250, NOW()),
  (uuid_generate_v4(), 'Ammo J', '6.5 Creedmoor', 'Standard 6.5 Creedmoor ammo', 'Rifle', 150, NOW()),
  (uuid_generate_v4(), 'Ammo K', '300 Win Mag', 'Standard 300 Win Mag ammo', 'Rifle', 100, NOW()),
  (uuid_generate_v4(), 'Ammo L', '50 BMG', 'Standard 50 BMG ammo', 'Rifle', 50, NOW());

-- Mock data for ammos_history table
INSERT INTO ammo_history (id, ammo_id, action, actor_id, action_timestamp) VALUES
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo A'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo B'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo C'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo D'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo E'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo F'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo G'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo H'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo I'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo J'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo K'), 'Created', uuid_generate_v4(), NOW()),
  (uuid_generate_v4(), (SELECT id FROM ammos WHERE name = 'Ammo L'), 'Created', uuid_generate_v4(), NOW());

COMMIT;
