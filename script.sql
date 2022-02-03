create table course
(
    CourseID   bigint auto_increment
        primary key,
    CourseName varchar(40) not null,
    Cap        int         not null,
    TeacherID  bigint      null,
    RestCap    int         null,
    constraint course_CourseID_uindex
        unique (CourseID)
);

create table courseStudent
(
    StudentID bigint not null
        primary key,
    Courses   text   null,
    constraint courseStudent_StudentID_uindex
        unique (StudentID)
);

create table user
(
    UserID    bigint auto_increment
        primary key,
    UserName  varchar(20)   not null,
    Password  varchar(20)   not null,
    UserState int default 0 not null,
    Nickname  varchar(20)   not null,
    UserType  int           not null,
    constraint user_UserID_uindex
        unique (UserID)
);


