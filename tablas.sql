create table USERS (
   ID               int              not null,
   RUN              int                  null,
   VD               varchar(50)          null,
   NAMES            varchar(50)          null,
   LAST_NAME        varchar(50)          null,
   GENDER           varchar(50)          null,
   BIRTHDAY         date                 null,
   PASS             varchar(50)          null,
   MAIL             varchar(50)          null,
   
   primary key (ID)
);

create table QUESTIONS (
   ID                 int              not null,
   QUESTION           varchar(50) not null,
   DESCRIPTION_Q      varchar(50)            null,
   DELETE_AT          bool     null,
   TEST_ID            int          null,
   
   primary key (ID)
)

create table ANSWERS (
   ID                 int              not null,
   QUESTION_ID        varchar(50)    not null,
   DESCRIPTION_Q      varchar(50)            null,
   DELETE_AT          bool     null,
   TEST_ID            int          null,

   primary key (ID)
)

create table TEST (
    ID                 int              not null,
    NAME_TEST varchar(50)    not null,
    CUT_POINT int null,
    MAX_POINT int null,
    OBSERVATION text null,

    primary key (ID)
)   