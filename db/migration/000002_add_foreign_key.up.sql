ALTER TABLE `course`      ADD CONSTRAINT `fk_course_category_id`    FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);
ALTER TABLE `transaction` ADD CONSTRAINT `fk_transaction_user_id`   FOREIGN KEY (`user_id`)     REFERENCES `user`     (`id`);
ALTER TABLE `transaction` ADD CONSTRAINT `fk_transaction_course_id` FOREIGN KEY (`course_id`)   REFERENCES `course`   (`id`);
