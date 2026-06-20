
CREATE TABLE public.subscriptions (
    id_sub bigint NOT NULL,
    service_name text,
    price_month integer,
    user_uuid uuid,
    start_date text NOT NULL,
    end_date text NOT NULL,
    CONSTRAINT subscriptions_end_date_check CHECK ((end_date ~ '^\d{2}-\d{4}$'::text)),
    CONSTRAINT subscriptions_start_date_check CHECK ((start_date ~ '^\d{2}-\d{4}$'::text))
);


CREATE SEQUENCE public.subscriptions_id_sub_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.subscriptions_id_sub_seq OWNED BY public.subscriptions.id_sub;

ALTER TABLE ONLY public.subscriptions ALTER COLUMN id_sub SET DEFAULT nextval('public.subscriptions_id_sub_seq'::regclass);
