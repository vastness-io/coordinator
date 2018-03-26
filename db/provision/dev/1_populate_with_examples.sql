INSERT INTO PROJECTS (NAME, TYPE) VALUES
  ('A', 'GITHUB'),
  ('B', 'BITBUCKET_SERVER');

INSERT INTO REPOSITORIES (ID, NAME) VALUES
  (1, 'repo-A'),
  (2, 'repo-B');

INSERT INTO BRANCHES (NAME, META, REPOSITORY_ID) VALUES
  ('branch-a', '{"languages" : {"Go" : 50, "Java": 50}}', 1),
  ('branch-b' ,'{"languages" : {"C" : 50, "Haskell": 50}}', 2);

INSERT INTO COMMITS (SHA, MESSAGE, TIMESTAMP, AUTHOR_NAME, AUTHOR_EMAIL, ADDED, MODIFIED, REMOVED) VALUES
  ('79511a5b6df54267d28ad9754c807ba1aebbb6e3', ' Add project service', now(), 'Joe Bloggs', 'joe.bloggs@gmail.com', '{}', '{}','{}');

INSERT INTO BRANCH_COMMITS (BRANCH_ID, COMMIT_SHA) VALUES
  ((SELECT b.ID FROM BRANCHES b WHERE  b.NAME = 'branch-a'), '79511a5b6df54267d28ad9754c807ba1aebbb6e3'),
  ((SELECT b.ID FROM BRANCHES b WHERE  b.NAME = 'branch-b'), '79511a5b6df54267d28ad9754c807ba1aebbb6e3');