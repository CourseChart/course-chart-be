package main

import "gorm.io/gorm"

func seed_activities(db *gorm.DB) {
	db.Exec(`
	INSERT INTO activities(id, name, description, metric, multiplier, custom)
	VALUES (1, 'Reading (understand)', '130 wpm; 10 pages per hour', '# of pages', 6, FALSE),
		(2, 'Reading (study guide)', '65 wpm; 5 pages per hour', '# of pages', 12, FALSE),
		(3, 'Writing (research)', '6 hours per page (500 words, single-spaced)', '# of pages', 360, FALSE),
		(4, 'Writing (reflection)', '90 minutes per page (500 words, single-spaced)', '# of pages', 90, FALSE),
		(5, 'Learning Objects (matching/multiple choice)', '10 minutes per object', '# of LOs', 10, FALSE),
		(6, 'Learning Objects (case study)', '20 minutes per object', '# of LOs', 20, FALSE),
		(7, 'Lecture', 'Factor 1.25x the actual lecture runtime', '# of minutes', 1.25, FALSE),
		(8, 'Videos', 'Factor the full length of video', '# of minutes', 1, FALSE),
		(9, 'Websites', '10-20 minutes', '# of minutes', 1, FALSE),
		(10, 'Discussion Boards', '250 words/60 minutes for initial post or 2 replies', '# of discussion boards', 60, FALSE),
		(11, 'Quizzes', 'Average 1.5 minutes per question', '# of questions', 1.5, FALSE),
		(12, 'Exams', 'Average 1.5 minutes per question', '# of questions', 1.5, FALSE),
		(13, 'Self Assessments', 'Average 1 minute per question', '# of questions', 1, FALSE),
		(14, 'Miscellaneous', 'any additional assignments not listed', '# of minutes', 1, FALSE);
		ALTER SEQUENCE activities_id_seq RESTART WITH 15;`)
}

func seed_full_course(db *gorm.DB) {
	db.Exec(`
		DELETE FROM module_activities;
		DELETE FROM modules;
		DELETE FROM courses;
		INSERT INTO courses(id, name, institution, credit_hours, length)
			VALUES (1, 'Foundations of Nursing', 'Colorado Nursing College', 3, 8);
		ALTER SEQUENCE courses_id_seq RESTART WITH 2;
		INSERT INTO modules(id, name, number, course_id)
			VALUES (1, 'Module 1', 1, 1),
			(2, 'Module 2', 2, 1),
			(3, 'Module 3', 3, 1),
			(4, 'Module 4', 4, 1),
			(5, 'Module 5', 5, 1),
			(6, 'Module 6', 6, 1),
			(7, 'Module 7', 7, 1),
			(8, 'Module 8', 8, 1);
		ALTER SEQUENCE modules_id_seq RESTART WITH 9;
		INSERT INTO module_activities(input, module_id, activity_id)
			VALUES 
			(107, 1, 1),
			(6, 1, 2),
			(7, 1, 5),
			(95, 1, 8),
			(1, 1, 10),
			(450, 1, 11),
			(50, 1, 13),
			( 53, 2, 1),
			( 5, 2, 2),
			( 5, 2, 5),
			( 71, 2, 8),
			( 1, 2, 10),
			( 100, 2, 11),
			( 66, 3, 1),
			( 4, 3, 2),
			( 1, 3, 4),
			( 4, 3, 5),
			( 2, 3, 6),
			( 86, 3, 8),
			( 1, 3, 10),
			( 240, 3, 11),
			( 50, 3, 13),
			( 105, 4, 1),
			( 7, 4, 2),
			( 2, 4, 4),
			( 3, 4, 5),
			( 75, 4, 8),
			( 390, 4, 11),
			( 50, 4, 13),
			( 52, 5, 1),
			( 5, 5, 2),
			( 1, 5, 4),
			( 5, 5, 5),
			( 1, 5, 6),
			( 62, 5, 8),
			( 1, 5, 10),
			( 300, 5, 11),
			( 36, 6, 1),
			( 5, 6, 2),
			( 5, 6, 5),
			( 1, 6, 6),
			( 40, 6, 8),
			( 1, 6, 10),
			( 90, 6, 11),
			( 50, 6, 13),
			( 88, 7, 1),
			( 5, 7, 2),
			( 4, 7, 5),
			( 2, 7, 6),
			( 42, 7, 8),
			( 240, 7, 11),
			( 3, 8, 3),
			( 100, 8, 13);
		ALTER SEQUENCE module_activities_id_seq RESTART WITH 54;
		`)
}
