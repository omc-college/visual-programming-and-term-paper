# Timetabling

### Разработать модуль управления расписанием и ресурсами



### Пользовательская история (Use Case)

##### Schedule lesson
AC: as a user I should be able to create a lesson, when Create lesson button was hit I see next form:
    * Teacher - (drop down list, with multiple selection, required)
    * Group - (drop down list, with multiple selection, required)
    * Room - (drop down list)
    * Start time and duration
* When user hits Save button entity validation should be performed.
* If validation successful lesson should be saved.
    * If there is an error, it should be shown for user.

##### Show Timetable
AC: as a user I should be able to get schedule for:
 * Group
 * Teacher
 * Room
I should be able to select schedule timelines:
 * Week
 
Timetable should be rendered as table where columns are dates and rows are lessons.

 
 
##### Reschedule lesson
AC: As a user I should be able to change any property of scheduled lesson if possible.
* When lesson object on timetable is hit, lesson edit window.
    * When user hits Save button entity validation is performed.
    * If validation successufull lesson should be saved.
        * If there is an error, it should be shown for user.
* If user hits cancel button, lesson edit window should be closed.

##### Delete lesson
AC: As a user I should be able to delete scheduled lesson.




##### main entities
 * CRUD 
 * POST api/v1/lesson/validate
 * timetable it is just view of lessons per properties
    - ~~time slot template (time from - time to) ?~~
    - group (external)
    - subject (external)
    - lesson (subject + group + time slot + lecturer(s) + room + other resources (projector, lab stuff))
 *Timetable lesson retrieve
     filters by group, teacher, resourse, timeline
 *Timetable lesson add
     - unique constrain 
         * room + timeslot + lecturer
 *Timetable lesson edit
 *Timetable lesson remove 
    
    
    			
###План работ
1. Design API using openAPI spec.
    * endpoints
    * entities (with relations)
2. start implementation from:
    * Schedule lesson
    * Show Timetable
    * ....