CREATE TABLE IF NOT EXISTS `users` (
  `id` VARCHAR (255) PRIMARY KEY UNIQUE NOT NULL,
  `email` VARCHAR (255) UNIQUE NOT NULL,
  `username` VARCHAR (255) UNIQUE NOT NULL,
  `password` VARCHAR (255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `bots` (
  `id` VARCHAR (255) PRIMARY KEY UNIQUE NOT NULL,
  `user_id` VARCHAR (255) NOT NULL REFERENCES `users`(`id`),
  `token` VARCHAR (255) NOT NULL,
  `is_verified` BOOLEAN DEFAULT FALSE NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `devel` BOOLEAN NOT NULL DEFAULT false,
  `version` VARCHAR (10) NOT NULL DEFAULT '',
  `handlegroups` BOOLEAN NOT NULL DEFAULT TRUE,
  `handlebroadcast` BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS `webhooks` (
  `context` VARCHAR (255) NOT NULL REFERENCES `bots`(`id`),
  `url` VARCHAR (255) NOT NULL,
  `forwardinternal` BIT NOT NULL DEFAULT 0 ,
  CONSTRAINT `webhooks_pkey` PRIMARY KEY (`context`, `url`)
);