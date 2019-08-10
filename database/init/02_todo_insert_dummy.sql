USE app1;
-- Insert initlal data
START TRANSACTION;
INSERT INTO todo(title, detail, auther) VALUES('First task', 'test data1', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Second task', 'test data2', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Third task', 'test data3', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data4', 'long long long long long', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data5', 'long long long long long', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data6', 'long long long long long', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data7', 'long long long long long', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data8', 'long long long long long', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data9', 'long long long long long', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data10', 'long long long long long', 'ketakom');
INSERT INTO todo(title, detail, auther) VALUES('Dummy Data11', 'long long long long long', 'ketakom');
COMMIT;