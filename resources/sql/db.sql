PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE configurations (id int primary key not null, cfgKey char(100) not null, cfgValue char(100) default null);
INSERT INTO "configurations" VALUES(1,'AppName','Gin-Skeleton');
INSERT INTO "configurations" VALUES(2,'AppVersion','0.0.1');
COMMIT;
