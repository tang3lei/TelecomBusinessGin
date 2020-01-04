create table employees
(
    id         int          not null identity (1,1),
    job_number nvarchar(32) not null,
    password   nvarchar(256) not null ,
    name       nvarchar(64) not null,
    create_at  datetime     not null DEFAULT CURRENT_TIMESTAMP,
    update_at  datetime     not null DEFAULT CURRENT_TIMESTAMP
        PRIMARY KEY (id)
);
go
create UNIQUE INDEX uniq_employee
    on employees (job_number, name)
GO

create TRIGGER employees_update_trigger
    on employees
    for update
    as
    update employees
    set update_at = getdate()
    where id in (select DISTINCT id from inserted);
go

ALTER TABLE employees
    ENABLE TRIGGER employees_update_trigger
go

