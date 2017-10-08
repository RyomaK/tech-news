CREATE TABLE `companies` (
	  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
		  `company` char(255) NOT NULL DEFAULT '',
			  PRIMARY KEY (`id`)
			) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;

CREATE TABLE `articles` (
	  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
		  `title` char(255) NOT NULL DEFAULT '',
			  `posted_at` date NOT NULL,
				  `url` char(255) NOT NULL DEFAULT '',
					  `text` char(255) DEFAULT '',
						  `company` char(255) NOT NULL DEFAULT '',
							  PRIMARY KEY (`id`)
							) ENGINE=InnoDB AUTO_INCREMENT=209 DEFAULT CHARSET=utf8;
