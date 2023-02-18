# dynpower-cli

Dynpower-CLI is a helper tool to manage the dynpower database

## Usage

It is recommended to use the dynpower Docker image. It comes with dynpower running as server and the dynpower CLI binary.

## Status

Currently dynpower is in development, there is no published version. When released, you'll find more information here.

```

CREATE TABLE `domains` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`domainname` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`access_key` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`dt_created` DATETIME NOT NULL,
	`dt_updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `domainname` (`domainname`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;

CREATE TABLE `dynrecords` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`domain_id` INT(11) NOT NULL DEFAULT '0',
	`hostname` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`dt_created` DATETIME NOT NULL,
	`dt_updated` DATETIME NOT NULL,
	`host_updated` DATETIME NOT NULL,
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `hostname` (`hostname`, `domain_id`) USING BTREE,
	INDEX `domain_id` (`domain_id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;
```
