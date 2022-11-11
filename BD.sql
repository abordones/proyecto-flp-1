/*==============================================================*/
/* Table: ANSWERS                                               */
/*==============================================================*/
create table ANSWERS 
(
   ID_A                 integer                        not null		AUTO_INCREMENT,
   ID_Q                 integer                        not null,
   POINT_A              integer                        null,
   OBSERVATION_A        varchar(200)                   null,
   ACTIVE_A             smallint                       null,
   constraint PK_ANSWERS primary key (ID_A)
);

/*==============================================================*/
/* Index: ANSWERS_PK                                            */
/*==============================================================*/
create unique index ANSWERS_PK on ANSWERS (
ID_A ASC
);

/*==============================================================*/
/* Index: TENER_FK                                              */
/*==============================================================*/
create index TENER_FK on ANSWERS (
ID_Q ASC
);

/*==============================================================*/
/* Table: PATIENTS                                              */
/*==============================================================*/
create table PATIENTS 
(
   ID_P                 integer                        not null		AUTO_INCREMENT,
   RUN_P                integer                        null,
   DV_P                 char(1)                        null,
   NAME_P               char(100)                      null,
   FATHERNAME_P         char(50)                       null,
   MOTHERNAME_P         char(50)                       null,
   BIRTHDAY_P           date                           null,
   PHONE_P              integer                        null,
   EMAIL_P              varchar(70)                    null,
   OBSERVATION_P        varchar(500)                   null,
   ACTIVE_P             smallint                       null,
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
   ID_PO                integer                        not null		AUTO_INCREMENT,
   ID_U                 integer                        not null,
   ID_T                 integer                        not null,
   constraint PK_POLLS primary key (ID_PO)
);

/*==============================================================*/
/* Index: POLLS_PK                                              */
/*==============================================================*/
create unique index POLLS_PK on POLLS (
ID_PO ASC
);

/*==============================================================*/
/* Index: REALIZAR_FK                                           */
/*==============================================================*/
create index REALIZAR_FK on POLLS (
ID_U ASC
);

/*==============================================================*/
/* Index: ESTAR_FK                                              */
/*==============================================================*/
create index ESTAR_FK on POLLS (
ID_T ASC
);

/*==============================================================*/
/* Table: QUESTIONS                                             */
/*==============================================================*/
create table QUESTIONS 
(
   ID_Q                 integer                        not null		AUTO_INCREMENT,
   ID_T                 integer                        not null,
   QUESTION_Q           varchar(150)                   null,
   DESCRIPTION_Q        varchar(200)                   null,
   ACTIVE_Q             smallint                       null,
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
/* Table: SESSIONS                                              */
/*==============================================================*/
create table SESSIONS 
(
   ID_P                 integer                        not null,
   ID_PO                integer                        not null,
   constraint PK_SESSIONS primary key (ID_P, ID_PO)
);

/*==============================================================*/
/* Index: SESSION_PK                                            */
/*==============================================================*/
create unique index SESSION_PK on SESSIONS (
ID_P ASC,
ID_PO ASC
);

/*==============================================================*/
/* Index: SESSION2_FK                                           */
/*==============================================================*/
create index SESSION2_FK on SESSIONS (
ID_PO ASC
);

/*==============================================================*/
/* Index: SESSION_FK                                            */
/*==============================================================*/
create index SESSION_FK on SESSIONS (
ID_P ASC
);

/*==============================================================*/
/* Table: TESTS                                                 */
/*==============================================================*/
create table TESTS 
(
   ID_T                 integer                        not null		AUTO_INCREMENT,
   NAME_T               char(100)                      null,
   CUTPOINT_T           integer                        null,
   MATCHPOINT_T         integer                        null,
   DESCRIPTION_T        varchar(500)                   null,
   ACTIVE_T             smallint                       null,
   constraint PK_TESTS primary key (ID_T)
);

/*==============================================================*/
/* Index: TESTS_PK                                              */
/*==============================================================*/
create unique index TESTS_PK on TESTS (
ID_T ASC
);

/*==============================================================*/
/* Table: USERS                                                 */
/*==============================================================*/
create table USERS 
(
   ID_U                 integer                        not null		AUTO_INCREMENT,
   RUN_U                integer                        null,
   DV_U                 char(1)                        null,
   NAME_U               char(100)                      null,
   FATHERNAME_U         char(50)                       null,
   MOTHERNAME_U         char(50)                       null,
   BIRTHDAY_U           date                           null,
   EMAIL_U              varchar(70)                    null,
   PASSWORD_U           varchar(70)                    null,
   ACTIVE_U             smallint                       null,
   constraint PK_USERS primary key (ID_U)
);

/*==============================================================*/
/* Index: USERS_PK                                              */
/*==============================================================*/
create unique index USERS_PK on USERS (
ID_U ASC
);

alter table ANSWERS
   add constraint FK_ANSWERS_TENER_QUESTION foreign key (ID_Q)
      references QUESTIONS (ID_Q)
      on update restrict
      on delete restrict;

alter table POLLS
   add constraint FK_POLLS_ESTAR_TESTS foreign key (ID_T)
      references TESTS (ID_T)
      on update restrict
      on delete restrict;

alter table POLLS
   add constraint FK_POLLS_REALIZAR_USERS foreign key (ID_U)
      references USERS (ID_U)
      on update restrict
      on delete restrict;

alter table QUESTIONS
   add constraint FK_QUESTION_PREGUNTAR_TESTS foreign key (ID_T)
      references TESTS (ID_T)
      on update restrict
      on delete restrict;

alter table SESSIONS
   add constraint FK_SESSIONS_SESSIONS_PATIENTS foreign key (ID_P)
      references PATIENTS (ID_P)
      on update restrict
      on delete restrict;

alter table SESSIONS
   add constraint FK_SESSIONS_SESSIONS2_POLLS foreign key (ID_PO)
      references POLLS (ID_PO)
      on update restrict
      on delete restrict;