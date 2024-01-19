--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 14.10 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: user_role; Type: TYPE; Schema: public; Owner: test
--

CREATE TYPE public.user_role AS ENUM (
    'admin',
    'student',
    'teacher'
);


ALTER TYPE public.user_role OWNER TO test;

--
-- Name: add_classgroup_to_current_schoolyear(); Type: FUNCTION; Schema: public; Owner: test
--

CREATE FUNCTION public.add_classgroup_to_current_schoolyear() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
    current_schoolyear_id INTEGER;
BEGIN
    -- Find the ID of the current SchoolYear
    SELECT id INTO current_schoolyear_id FROM school_years WHERE is_current = TRUE LIMIT 1;

    -- If a current SchoolYear exists, create a relation in the join table
    IF FOUND THEN
        INSERT INTO school_year_class_group (school_year_id, class_group_id) VALUES (current_schoolyear_id, NEW.id);
    END IF;

    -- Return the new row
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.add_classgroup_to_current_schoolyear() OWNER TO test;

--
-- Name: open_new_school_year(text, timestamp without time zone, timestamp without time zone); Type: FUNCTION; Schema: public; Owner: test
--

CREATE FUNCTION public.open_new_school_year(year_name text, start_date timestamp without time zone, end_date timestamp without time zone) RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
    new_year_id INT;
BEGIN
    -- Check if school year exists
    IF NOT EXISTS (SELECT 1 FROM school_years WHERE name = year_name) THEN
        -- Make all previous years archived
        UPDATE school_years SET is_current = false;
        -- Add new school year
        INSERT INTO school_years ("name", "start", "end", is_current, created_at, updated_at) VALUES (year_name, start_date, end_date, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id INTO new_year_id;

        -- Insert records into school_year_class_group for class groups with education_year < 8
        INSERT INTO school_year_class_group (school_year_id, class_group_id)
        SELECT new_year_id, id FROM class_groups WHERE education_year < 8;

        -- Increment education_year by one for each class group
        UPDATE class_groups SET education_year = education_year + 1 WHERE education_year < 8;

    ELSE
        RAISE EXCEPTION 'School year with this name exists';
    END IF;
END;
$$;


ALTER FUNCTION public.open_new_school_year(year_name text, start_date timestamp without time zone, end_date timestamp without time zone) OWNER TO test;

--
-- Name: set_mark_school_year(); Type: FUNCTION; Schema: public; Owner: test
--

CREATE FUNCTION public.set_mark_school_year() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW."school_year_id" IS NULL THEN
        -- Pobierz ID roku szkolnego, który jest oznaczony jako aktualny
        SELECT id INTO NEW."school_year_id"
        FROM school_years
        WHERE is_current = true;
    END IF;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_mark_school_year() OWNER TO test;

--
-- Name: set_school_year(); Type: FUNCTION; Schema: public; Owner: test
--

CREATE FUNCTION public.set_school_year() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW."school_year_id" IS NULL THEN
        -- Pobierz ID roku szkolnego, który jest oznaczony jako aktualny
        SELECT id INTO NEW."school_year_id"
        FROM school_years
        WHERE is_current = true;
    END IF;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_school_year() OWNER TO test;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: class_groups; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.class_groups (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    education_year bigint DEFAULT 1
);


ALTER TABLE public.class_groups OWNER TO test;

--
-- Name: class_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.class_groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.class_groups_id_seq OWNER TO test;

--
-- Name: class_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.class_groups_id_seq OWNED BY public.class_groups.id;


--
-- Name: lessons; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.lessons (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    subject_id bigint,
    teacher_id bigint,
    class_group_id bigint,
    start timestamp with time zone,
    "end" timestamp with time zone,
    day_of_week bigint,
    school_year_id bigint
);


ALTER TABLE public.lessons OWNER TO test;

--
-- Name: lessons_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.lessons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lessons_id_seq OWNER TO test;

--
-- Name: lessons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.lessons_id_seq OWNED BY public.lessons.id;


--
-- Name: marks; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.marks (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    value numeric,
    student_id bigint,
    comment text,
    date timestamp with time zone,
    subject_id bigint,
    teacher_id bigint,
    school_year_id bigint
);


ALTER TABLE public.marks OWNER TO test;

--
-- Name: marks_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.marks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.marks_id_seq OWNER TO test;

--
-- Name: marks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.marks_id_seq OWNED BY public.marks.id;


--
-- Name: notifications; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.notifications (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    message text,
    read boolean
);


ALTER TABLE public.notifications OWNER TO test;

--
-- Name: notifications_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.notifications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notifications_id_seq OWNER TO test;

--
-- Name: notifications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.notifications_id_seq OWNED BY public.notifications.id;


--
-- Name: school_year_class_group; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.school_year_class_group (
    school_year_id bigint NOT NULL,
    class_group_id bigint NOT NULL
);


ALTER TABLE public.school_year_class_group OWNER TO test;

--
-- Name: school_years; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.school_years (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    is_current boolean,
    start timestamp with time zone,
    "end" timestamp with time zone
);


ALTER TABLE public.school_years OWNER TO test;

--
-- Name: school_years_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.school_years_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.school_years_id_seq OWNER TO test;

--
-- Name: school_years_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.school_years_id_seq OWNED BY public.school_years.id;


--
-- Name: students; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.students (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    class_group_id bigint
);


ALTER TABLE public.students OWNER TO test;

--
-- Name: students_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.students_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.students_id_seq OWNER TO test;

--
-- Name: students_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.students_id_seq OWNED BY public.students.id;


--
-- Name: subject_teacher_classes; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.subject_teacher_classes (
    teacher_id bigint,
    subject_id bigint,
    class_group_id bigint
);


ALTER TABLE public.subject_teacher_classes OWNER TO test;

--
-- Name: subject_teachers; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.subject_teachers (
    teacher_id bigint NOT NULL,
    subject_id bigint NOT NULL
);


ALTER TABLE public.subject_teachers OWNER TO test;

--
-- Name: subjects; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.subjects (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text
);


ALTER TABLE public.subjects OWNER TO test;

--
-- Name: subjects_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.subjects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subjects_id_seq OWNER TO test;

--
-- Name: subjects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.subjects_id_seq OWNED BY public.subjects.id;


--
-- Name: teachers; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.teachers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint
);


ALTER TABLE public.teachers OWNER TO test;

--
-- Name: teachers_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.teachers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.teachers_id_seq OWNER TO test;

--
-- Name: teachers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.teachers_id_seq OWNED BY public.teachers.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    surname text,
    email text,
    role public.user_role,
    password text,
    active boolean
);


ALTER TABLE public.users OWNER TO test;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: test
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO test;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: test
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: class_groups id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.class_groups ALTER COLUMN id SET DEFAULT nextval('public.class_groups_id_seq'::regclass);


--
-- Name: lessons id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.lessons ALTER COLUMN id SET DEFAULT nextval('public.lessons_id_seq'::regclass);


--
-- Name: marks id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.marks ALTER COLUMN id SET DEFAULT nextval('public.marks_id_seq'::regclass);


--
-- Name: notifications id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.notifications ALTER COLUMN id SET DEFAULT nextval('public.notifications_id_seq'::regclass);


--
-- Name: school_years id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.school_years ALTER COLUMN id SET DEFAULT nextval('public.school_years_id_seq'::regclass);


--
-- Name: students id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.students ALTER COLUMN id SET DEFAULT nextval('public.students_id_seq'::regclass);


--
-- Name: subjects id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subjects ALTER COLUMN id SET DEFAULT nextval('public.subjects_id_seq'::regclass);


--
-- Name: teachers id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.teachers ALTER COLUMN id SET DEFAULT nextval('public.teachers_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: class_groups; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.class_groups (id, created_at, updated_at, deleted_at, name, education_year) FROM stdin;
1	2024-01-19 12:52:52.003878+00	2024-01-19 12:52:52.003878+00	\N	1 mat-fiz	1
\.


--
-- Data for Name: lessons; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.lessons (id, created_at, updated_at, deleted_at, subject_id, teacher_id, class_group_id, start, "end", day_of_week, school_year_id) FROM stdin;
1	2024-01-19 12:52:52.100813+00	2024-01-19 12:52:52.100813+00	\N	1	1	1	2024-01-19 08:00:00+00	2024-01-19 08:45:00+00	1	1
\.


--
-- Data for Name: marks; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.marks (id, created_at, updated_at, deleted_at, value, student_id, comment, date, subject_id, teacher_id, school_year_id) FROM stdin;
1	2024-01-19 12:52:52.085452+00	2024-01-19 12:52:52.085452+00	\N	3.5	1		\N	1	1	1
\.


--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.notifications (id, created_at, updated_at, deleted_at, user_id, message, read) FROM stdin;
\.


--
-- Data for Name: school_year_class_group; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.school_year_class_group (school_year_id, class_group_id) FROM stdin;
1	1
\.


--
-- Data for Name: school_years; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.school_years (id, created_at, updated_at, deleted_at, name, is_current, start, "end") FROM stdin;
1	2024-01-19 12:52:51.997824+00	2024-01-19 12:52:51.997824+00	\N	2023/2024	t	2023-09-01 00:00:00+00	2024-06-30 00:00:00+00
\.


--
-- Data for Name: students; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.students (id, created_at, updated_at, deleted_at, user_id, class_group_id) FROM stdin;
1	2024-01-19 12:52:52.074448+00	2024-01-19 12:52:52.074448+00	\N	3	1
\.


--
-- Data for Name: subject_teacher_classes; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.subject_teacher_classes (teacher_id, subject_id, class_group_id) FROM stdin;
1	1	1
\.


--
-- Data for Name: subject_teachers; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.subject_teachers (teacher_id, subject_id) FROM stdin;
1	1
\.


--
-- Data for Name: subjects; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.subjects (id, created_at, updated_at, deleted_at, name) FROM stdin;
1	2024-01-19 12:52:51.989326+00	2024-01-19 12:52:51.989326+00	\N	Wychowanie Fizyczne
\.


--
-- Data for Name: teachers; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.teachers (id, created_at, updated_at, deleted_at, user_id) FROM stdin;
1	2024-01-19 12:52:51.954354+00	2024-01-19 12:52:51.954354+00	\N	2
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: test
--

COPY public.users (id, created_at, updated_at, deleted_at, name, surname, email, role, password, active) FROM stdin;
1	2024-01-19 12:52:51.892327+00	2024-01-19 12:52:51.892327+00	\N	Marcin	Testowy	admin@e-student.com	admin	$2a$10$h/AHHs6en0sAt0cmxvFlHOmE7C28ZGd66J6IBFQNNesC382Ix5wki	t
2	2024-01-19 12:52:51.953459+00	2024-01-19 12:52:51.953459+00	\N	Adam	Kowalski	teacher@e-student.com	teacher	$2a$10$Ga6tSbl4q8e/O2gH1zvjDOJ.6JeGxORzCoe9lBz4OmO15WzcMx/fq	t
3	2024-01-19 12:52:52.072548+00	2024-01-19 12:52:52.072548+00	\N	Maciej	Szkolny	student@e-student.com	student	$2a$10$aOwSSpHFs0sDJWoTfmKLsugZu2Cx5EV5svq/yZ1Y0lg8S6KsuTqZ6	t
\.


--
-- Name: class_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.class_groups_id_seq', 1, true);


--
-- Name: lessons_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.lessons_id_seq', 1, true);


--
-- Name: marks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.marks_id_seq', 1, true);


--
-- Name: notifications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.notifications_id_seq', 1, false);


--
-- Name: school_years_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.school_years_id_seq', 1, true);


--
-- Name: students_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.students_id_seq', 1, true);


--
-- Name: subjects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.subjects_id_seq', 1, true);


--
-- Name: teachers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.teachers_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: test
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


--
-- Name: class_groups class_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.class_groups
    ADD CONSTRAINT class_groups_pkey PRIMARY KEY (id);


--
-- Name: lessons lessons_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT lessons_pkey PRIMARY KEY (id);


--
-- Name: marks marks_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.marks
    ADD CONSTRAINT marks_pkey PRIMARY KEY (id);


--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: school_year_class_group school_year_class_group_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.school_year_class_group
    ADD CONSTRAINT school_year_class_group_pkey PRIMARY KEY (school_year_id, class_group_id);


--
-- Name: school_years school_years_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.school_years
    ADD CONSTRAINT school_years_pkey PRIMARY KEY (id);


--
-- Name: students students_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);


--
-- Name: subject_teachers subject_teachers_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subject_teachers
    ADD CONSTRAINT subject_teachers_pkey PRIMARY KEY (teacher_id, subject_id);


--
-- Name: subjects subjects_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subjects
    ADD CONSTRAINT subjects_pkey PRIMARY KEY (id);


--
-- Name: teachers teachers_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT teachers_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_class_groups_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_class_groups_deleted_at ON public.class_groups USING btree (deleted_at);


--
-- Name: idx_lessons_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_lessons_deleted_at ON public.lessons USING btree (deleted_at);


--
-- Name: idx_marks_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_marks_deleted_at ON public.marks USING btree (deleted_at);


--
-- Name: idx_notifications_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_notifications_deleted_at ON public.notifications USING btree (deleted_at);


--
-- Name: idx_school_years_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_school_years_deleted_at ON public.school_years USING btree (deleted_at);


--
-- Name: idx_students_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_students_deleted_at ON public.students USING btree (deleted_at);


--
-- Name: idx_subjects_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_subjects_deleted_at ON public.subjects USING btree (deleted_at);


--
-- Name: idx_teachers_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_teachers_deleted_at ON public.teachers USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: test
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: class_groups add_classgroup_after_insert; Type: TRIGGER; Schema: public; Owner: test
--

CREATE TRIGGER add_classgroup_after_insert AFTER INSERT ON public.class_groups FOR EACH ROW EXECUTE FUNCTION public.add_classgroup_to_current_schoolyear();


--
-- Name: lessons trigger_set_lesson_school_year; Type: TRIGGER; Schema: public; Owner: test
--

CREATE TRIGGER trigger_set_lesson_school_year BEFORE INSERT ON public.lessons FOR EACH ROW EXECUTE FUNCTION public.set_school_year();


--
-- Name: marks trigger_set_mark_school_year; Type: TRIGGER; Schema: public; Owner: test
--

CREATE TRIGGER trigger_set_mark_school_year BEFORE INSERT ON public.marks FOR EACH ROW EXECUTE FUNCTION public.set_school_year();


--
-- Name: students fk_class_groups_students; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT fk_class_groups_students FOREIGN KEY (class_group_id) REFERENCES public.class_groups(id);


--
-- Name: lessons fk_lessons_class_group; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT fk_lessons_class_group FOREIGN KEY (class_group_id) REFERENCES public.class_groups(id);


--
-- Name: lessons fk_lessons_school_year; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT fk_lessons_school_year FOREIGN KEY (school_year_id) REFERENCES public.school_years(id);


--
-- Name: lessons fk_lessons_subject; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT fk_lessons_subject FOREIGN KEY (subject_id) REFERENCES public.subjects(id);


--
-- Name: lessons fk_lessons_teacher; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT fk_lessons_teacher FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


--
-- Name: marks fk_marks_school_year; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.marks
    ADD CONSTRAINT fk_marks_school_year FOREIGN KEY (school_year_id) REFERENCES public.school_years(id);


--
-- Name: notifications fk_notifications_user; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notifications_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: school_year_class_group fk_school_year_class_group_class_group; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.school_year_class_group
    ADD CONSTRAINT fk_school_year_class_group_class_group FOREIGN KEY (class_group_id) REFERENCES public.class_groups(id);


--
-- Name: school_year_class_group fk_school_year_class_group_school_year; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.school_year_class_group
    ADD CONSTRAINT fk_school_year_class_group_school_year FOREIGN KEY (school_year_id) REFERENCES public.school_years(id);


--
-- Name: marks fk_students_marks; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.marks
    ADD CONSTRAINT fk_students_marks FOREIGN KEY (student_id) REFERENCES public.students(id);


--
-- Name: students fk_students_user; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT fk_students_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: subject_teacher_classes fk_subject_teacher_classes_class_group; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subject_teacher_classes
    ADD CONSTRAINT fk_subject_teacher_classes_class_group FOREIGN KEY (class_group_id) REFERENCES public.class_groups(id);


--
-- Name: subject_teacher_classes fk_subject_teacher_classes_subject; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subject_teacher_classes
    ADD CONSTRAINT fk_subject_teacher_classes_subject FOREIGN KEY (subject_id) REFERENCES public.subjects(id);


--
-- Name: subject_teacher_classes fk_subject_teacher_classes_teacher; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subject_teacher_classes
    ADD CONSTRAINT fk_subject_teacher_classes_teacher FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


--
-- Name: subject_teachers fk_subject_teachers_subject; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subject_teachers
    ADD CONSTRAINT fk_subject_teachers_subject FOREIGN KEY (subject_id) REFERENCES public.subjects(id);


--
-- Name: subject_teachers fk_subject_teachers_teacher; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.subject_teachers
    ADD CONSTRAINT fk_subject_teachers_teacher FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


--
-- Name: marks fk_subjects_marks; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.marks
    ADD CONSTRAINT fk_subjects_marks FOREIGN KEY (subject_id) REFERENCES public.subjects(id);


--
-- Name: marks fk_teachers_marks; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.marks
    ADD CONSTRAINT fk_teachers_marks FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


--
-- Name: teachers fk_teachers_user; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT fk_teachers_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

