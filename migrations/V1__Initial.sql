--
-- Name: ratings; Type: TABLE; Schema: public; Owner: goddd
--

CREATE TABLE IF NOT EXISTS ratings (
                         id uuid NOT NULL,
                         recipe_id uuid NOT NULL,
                         value smallint NOT NULL,
                         PRIMARY KEY (id)
);
