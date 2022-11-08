/*==============================================================*/
/* Table: ANSWERS                                               */
/*==============================================================*/
create table ANSWERS 
(
   ID_A                 integer                        not null		AUTO_INCREMENT,
   ID_T                 integer                        not null,
   POINT_A              integer                        null,
   OBSERVATION_A        varchar(200)                   null,
   ACTIVE_A             smallint                       null,
   constraint PK_ANSWERS primary key (ID_A)
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
   FATHERNAME_P         char(100)                      null,
   MOTHERNAME_P         char(100)                      null,
   PHONE_P              integer                        null,
   EMAIL_P              varchar(100)                   null,
   BIRTHDAY_P           date                           null,
   OBSERVATION_P        varchar(200)                   null,
   ACTIVE_P             smallint                       null,
   constraint PK_PATIENTS primary key (ID_P)
);

/*==============================================================*/
/* Table: POLLS                                                 */
/*==============================================================*/
create table POLLS 
(
   ID_PO                integer                        not null		AUTO_INCREMENT,
   ID_T                 integer                        not null,
   ID_U                 integer                        not null,
   constraint PK_POLLS primary key (ID_PO)
);

/*==============================================================*/
/* Table: QUESTIONS                                             */
/*==============================================================*/
create table QUESTIONS 
(
   ID_Q                 integer                        not null		AUTO_INCREMENT,
   ID_T                 integer                        not null,
   QUESTION_Q           varchar(100)                   null,
   DESCRIPTION_Q        varchar(200)                   null,
   ACTIVE_Q             smallint                       null,
   constraint PK_QUESTIONS primary key (ID_Q)
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
   OBSERVATION_T        varchar(200)                   null,
   ACTIVE_T             smallint                       null,
   constraint PK_TESTS primary key (ID_T)
);

/*==============================================================*/
/* Table: USAR                                                  */
/*==============================================================*/
create table USAR 
(
   ID_PO                integer                        not null,
   ID_P                 integer                        not null,
   DATE                 date                           null,
   constraint PK_USAR primary key (ID_PO, ID_P)
);

/*==============================================================*/
/* Table: USERS                                                 */
/*==============================================================*/
create table USERS 
(
   ID_U                 integer                        not null		AUTO_INCREMENT,
   RUN_U                integer                        null,
   DV_U                 char(1)                        null,
   NAME_U               char(200)                      null,
   FATHERNAME_U         char(100)                      null,
   MOTHERNAME_U         char(100)                      null,
   BIRTHDAY_U           date                           null,
   PASSWORD_U           varchar(100)                   null,
   EMAIL_U              varchar(100)                   null,
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
   add constraint FK_ANSWERS_RESPONDER_TESTS foreign key (ID_T)
      references TESTS (ID_T)
      on update restrict
      on delete restrict;

alter table POLLS
   add constraint FK_POLLS_ESTAR_TESTS foreign key (ID_T)
      references TESTS (ID_T)
      on update restrict
      on delete restrict;

alter table POLLS
   add constraint FK_POLLS_HACER_USERS foreign key (ID_U)
      references USERS (ID_U)
      on update restrict
      on delete restrict;

alter table QUESTIONS
   add constraint FK_QUESTION_PREGUNTAR_TESTS foreign key (ID_T)
      references TESTS (ID_T)
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
