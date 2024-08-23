CREATE TABLE public.player (
	smartid uuid NOT NULL,
	firstname varchar NOT NULL,
	lastname varchar NOT NULL,
	"position" varchar NOT NULL,
	currentteamid int4 NOT NULL,
	CONSTRAINT player_pk PRIMARY KEY (smartid)
);

CREATE TABLE public."passing" (
	id int4 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE) NOT NULL,
	season int4 NOT NULL,
	week int4 NOT NULL,
	seasontype varchar NOT NULL,
	completions int4 NOT NULL,
	attempts int4 NOT NULL,
	interceptions int4 NOT NULL,
	passtouchdowns int4 NOT NULL,
	playerid uuid NOT NULL,
	CONSTRAINT passing_pk PRIMARY KEY (id),
	CONSTRAINT passing_player_fk FOREIGN KEY (playerid) REFERENCES public.player(smartid)
);
