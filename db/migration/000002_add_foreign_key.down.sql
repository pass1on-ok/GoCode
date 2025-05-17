ALTER TABLE `course`         DROP FOREIGN KEY    `fk_course_category_id`;
ALTER TABLE `transaction`    DROP FOREIGN KEY    `fk_transaction_user_id`;
ALTER TABLE `transaction`    DROP FOREIGN KEY    `fk_transaction_course_id`;
