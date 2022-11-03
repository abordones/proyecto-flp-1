/*==============================================================*/
/* DBMS name:      Sybase SQL Anywhere 12                       */
/* Created on:     03-11-2022 12:49:52                          */
/*==============================================================*/


if exists(select 1 from sys.sysforeignkey where role='FK_ANSWERS_RESPONDER_TESTS') then
    alter table ANSWERS
       delete foreign key FK_ANSWERS_RESPONDER_TESTS
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_POLLS_EJERCER_USERS') then
    alter table POLLS
       delete foreign key FK_POLLS_EJERCER_USERS
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_QUESTION_PREGUNTAR_TESTS') then
    alter table QUESTIONS
       delete foreign key FK_QUESTION_PREGUNTAR_TESTS
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_TESTS_ESTAR_POLLS') then
    alter table TESTS
       delete foreign key FK_TESTS_ESTAR_POLLS
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_USAR_USAR_POLLS') then
    alter table USAR
       delete foreign key FK_USAR_USAR_POLLS
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_USAR_USAR2_PATIENTS') then
    alter table USAR
       delete foreign key FK_USAR_USAR2_PATIENTS
end if;

drop index if exists ANSWERS.RESPONDER_FK;

drop index if exists ANSWERS.ANSWERS_PK;

drop table if exists ANSWERS;

drop index if exists PATIENTS.PATIENTS_PK;

drop table if exists PATIENTS;

drop index if exists POLLS.EJERCER_FK;

drop index if exists POLLS.POLLS_PK;

drop table if exists POLLS;

drop index if exists QUESTIONS.PREGUNTAR_FK;

drop index if exists QUESTIONS.QUESTIONS_PK;

drop table if exists QUESTIONS;

drop index if exists TESTS.ESTAR_FK;

drop index if exists TESTS.TESTS_PK;

drop table if exists TESTS;

drop index if exists USAR.USAR_FK;

drop index if exists USAR.USAR2_FK;

drop index if exists USAR.USAR_PK;

drop table if exists USAR;

drop index if exists USERS.USERS_PK;

drop table if exists USERS;

/*==============================================================*/
/* Table: ANSWERS                                               */
/*==============================================================*/
create table ANSWERS 
(
   ID_A                 integer                        not null,
   ID_T                 integer                        not null,
   POINT_A              integer                        null,
   OBSERVATION_A        long varchar                   null,
   constraint PK_ANSWERS primary key (ID_A)
);

/*==============================================================*/
/* Index: ANSWERS_PK                                            */
/*==============================================================*/
create unique index ANSWERS_PK on ANSWERS (
ID_A ASC
);

/*==============================================================*/
/* Index: RESPONDER_FK                                          */
/*==============================================================*/
create index RESPONDER_FK on ANSWERS (
ID_T ASC
);

/*==============================================================*/
/* Table: PATIENTS                                              */
/*==============================================================*/
create table PATIENTS 
(
   ID_P                 integer                        not null,
   RUN_P                integer                        null,
   DV_P                 char(1)                        null,
   NAME_P               char(100)                      null,
   FATHERNAME_P         char(100)                      null,
   MOTHERNAME_P         char(100)                      null,
   PHONE_P              integer                        null,
   EMAIL_P              varchar(100)                   null,
   OBSERVATION_P        long varchar                   null,
   DELETE_P             smallint                       null,
   constraint PK_PATIENTS primary key (ID_P)
);

/*==============================================================*/
/* Index: PATIENTS_PK                                           */
/*==============================================================*/
create unique index PATIENTS_PK on PATIENTS (
ID_P ASC
);

/*==============================================================*/
/* Table: POLLS                                                 */
/*==============================================================*/
create table POLLS 
(
   ID_PO                integer                        not null,
   ID                   char(10)                       not null,
   constraint PK_POLLS primary key (ID_PO)
);

/*==============================================================*/
/* Index: POLLS_PK                                              */
/*==============================================================*/
create unique index POLLS_PK on POLLS (
ID_PO ASC
);

/*==============================================================*/
/* Index: EJERCER_FK                                            */
/*==============================================================*/
create index EJERCER_FK on POLLS (
ID ASC
);

/*==============================================================*/
/* Table: QUESTIONS                                             */
/*==============================================================*/
create table QUESTIONS 
(
   ID_Q                 integer                        not null,
   ID_T                 integer                        not null,
   QUESTION_Q           varchar(200)                   null,
   DESCRIPTION_Q        long varchar                   null,
   DELETE_Q             smallint                       null,
   constraint PK_QUESTIONS primary key (ID_Q)
);

/*==============================================================*/
/* Index: QUESTIONS_PK                                          */
/*==============================================================*/
create unique index QUESTIONS_PK on QUESTIONS (
ID_Q ASC
);

/*==============================================================*/
/* Index: PREGUNTAR_FK                                          */
/*==============================================================*/
create index PREGUNTAR_FK on QUESTIONS (
ID_T ASC
);

/*==============================================================*/
/* Table: TESTS                                                 */
/*==============================================================*/
create table TESTS 
(
   ID_T                 integer                        not null,
   ID_PO                integer                        not null,
   NAME_T               char(100)                      null,
   CUTPOINT_T           integer                        null,
   MATCHPOINT_T         integer                        null,
   OBSERVATION_T        long varchar                   null,
   constraint PK_TESTS primary key (ID_T)
);

/*==============================================================*/
/* Index: TESTS_PK                                              */
/*==============================================================*/
create unique index TESTS_PK on TESTS (
ID_T ASC
);

/*==============================================================*/
/* Index: ESTAR_FK                                              */
/*==============================================================*/
create index ESTAR_FK on TESTS (
ID_PO ASC
);

/*==============================================================*/
/* Table: USAR                                                  */
/*==============================================================*/
create table USAR 
(
   ID_PO                integer                        not null,
   ID_P                 integer                        not null,
   constraint PK_USAR primary key (ID_PO, ID_P)
);

/*==============================================================*/
/* Index: USAR_PK                                               */
/*==============================================================*/
create unique index USAR_PK on USAR (
ID_PO ASC,
ID_P ASC
);

/*==============================================================*/
/* Index: USAR2_FK                                              */
/*==============================================================*/
create index USAR2_FK on USAR (
ID_P ASC
);

/*==============================================================*/
/* Index: USAR_FK                                               */
/*==============================================================*/
create index USAR_FK on USAR (
ID_PO ASC
);

/*==============================================================*/
/* Table: USERS                                                 */
/*==============================================================*/
create table USERS 
(
   ID                   char(10)                       not null,
   RUN                  integer                        null,
   DV                   char(1)                        null,
   NAMES                char(100)                      null,
   FATHERNAME           char(100)                      null,
   MOTHERNAME           char(100)                      null,
   DAY                  integer                        null,
   MONTH                integer                        null,
   YEAR                 integer                        null,
   PASSWORD             varchar(100)                   null,
   "DELETE"             smallint                       null,
   constraint PK_USERS primary key (ID)
);

/*==============================================================*/
/* Index: USERS_PK                                              */
/*==============================================================*/
create unique index USERS_PK on USERS (
ID ASC
);

alter table ANSWERS
   add constraint FK_ANSWERS_RESPONDER_TESTS foreign key (ID_T)
      references TESTS (ID_T)
      on update restrict
      on delete restrict;

alter table POLLS
   add constraint FK_POLLS_EJERCER_USERS foreign key (ID)
      references USERS (ID)
      on update restrict
      on delete restrict;

alter table QUESTIONS
   add constraint FK_QUESTION_PREGUNTAR_TESTS foreign key (ID_T)
      references TESTS (ID_T)
      on update restrict
      on delete restrict;

alter table TESTS
   add constraint FK_TESTS_ESTAR_POLLS foreign key (ID_PO)
      references POLLS (ID_PO)
      on update restrict
      on delete restrict;

alter table USAR
   add constraint FK_USAR_USAR_POLLS foreign key (ID_PO)
      references POLLS (ID_PO)
      on update restrict
      on delete restrict;

alter table USAR
   add constraint FK_USAR_USAR2_PATIENTS foreign key (ID_P)
      references PATIENTS (ID_P)
      on update restrict
      on delete restrict;

