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

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.apods (
    title text NOT NULL,
    explanation text NOT NULL,
    image text NOT NULL,
    media_type text NOT NULL,
    service_version text NOT NULL,
    date timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

ALTER TABLE public.apods OWNER TO postgres;

COPY public.apods (title, explanation, image, media_type, service_version, date, created_at) FROM stdin;
NGC 1097: Spiral Galaxy with Supernova	What's happening in the lower arm of this spiral galaxy? A supernova. Last month, supernova SN 2023rve was discovered with UAE's Al-Khatim Observatory and later found to be consistent with the death explosion of a massive star, possibly leaving behind a black hole. Spiral galaxy NGC 1097 is a relatively close 45 million light years away and visible with a small telescope toward the southern constellation of the Furnace (Fornax).  The galaxy is notable not only for its picturesque spiral arms, but also for faint jets consistent with ancient star streams left over from a galactic collision -- possibly with the small galaxy seen between its arms on the lower left. The featured image highlights the new supernova by blinking between two exposures taken several months apart. Finding supernovas in nearby galaxies can be important in determining the scale and expansion rate of our entire universe -- a topic currently of unexpected tension and much debate.    APOD editor to speak: in Houghton, Michigan on Thursday, October 12 at 6 pm	c44b5e15-c346-4085-a75c-c35385d74bcb.jpg	image	v1	2023-10-11 00:00:00	2023-10-11 12:34:25.40087
Hidden Orion from Webb	The Great Nebula in Orion has hidden stars.   To the unaided eye in visible light, it appears as a small fuzzy patch in the constellation of Orion. But this image was taken by the Webb Space Telescope in a representative-color composite of red and very near infrared light.  It confirms with impressive detail that the  Orion Nebula is a busy neighborhood of young stars, hot gas, and dark dust.  The rollover image shows the same image in representative colors further into the near infrared. The power behind much of the Orion Nebula (M42) is the Trapezium - a cluster of bright stars near the nebula's center.  The diffuse and filamentary glow surrounding the bright stars is mostly heated interstellar dust.  Detailed inspection of these images shows an unexpectedly large number of Jupiter-Mass Binary Objects (JuMBOs), pairs of Jupiter-mass objects which might give a clue to how stars are forming.  The whole Orion Nebula cloud complex, which includes the Horsehead Nebula, will slowly disperse over the next few million years.   APOD editor to speak: in Houghton, Michigan on Thursday, October 12 at 6 pm	4dc13999-f3f1-43dd-8b7c-e3d184ba9889.jpg	image	v1	2023-10-10 00:00:00	2023-10-11 12:34:26.840602
A Distorted Sunrise Eclipse	Yes, but have you ever seen a sunrise like this?  Here, after initial cloudiness, the Sun appeared to rise in two pieces and during a partial eclipse in 2019, causing the photographer to describe it as the most stunning sunrise of his life.  The dark circle near the top of the atmospherically-reddened Sun is the Moon -- but so is the dark peak just below it.  This is because along the way, the Earth's atmosphere had a layer of unusually warm air over the sea which acted like a gigantic lens and created a second image.  For a normal sunrise or sunset, this rare phenomenon of atmospheric optics is known as the Etruscan vase effect. The featured picture was captured in December 2019 from Al Wakrah, Qatar.  Some observers in a narrow band of Earth to the east were able to see a full annular solar eclipse -- where the Moon appears completely surrounded by the background Sun in a ring of fire.  The next solar eclipse, also an annular eclipse for well-placed observers, will occur this coming Saturday.   APOD editor to speak: in Houghton, Michigan on Thursday, October 12 at 6 pm	6c67494b-94b6-45d4-bfc0-58ccc381b855.jpg	image	v1	2023-10-09 00:00:00	2023-10-11 12:34:27.365351
Plane, Clouds, Moon, Spots, Sun	What's that in front of the Sun?  The closest object is an airplane, visible just below the Sun's center and caught purely by chance.  Next out are numerous clouds in Earth's atmosphere, creating a series of darkened horizontal streaks. Farther out is Earth's Moon, seen as the large dark circular bite on the upper right. Just above the airplane and just below the Sun's surface are sunspots. The main sunspot group captured here, AR 2192, was in 2014 one of the largest ever recorded and had been crackling and bursting with flares since it came around the edge of the Sun a week before. This show of solar silhouettes was unfortunately short-lived.  Within a few seconds the plane flew away. Within a few minutes the clouds drifted off. Within a few hours the partial solar eclipse of the Sun by the Moon was over. Fortunately, when it comes to the Sun, even unexpected  alignments are surprisingly frequent. Perhaps one will be imaged this Saturday when a new partial solar eclipse will be visible from much of North and South America.    APOD editor to speak: in Houghton, Michigan on Thursday, October 12 at 6 pm	f045af25-55e1-4256-8d4c-1fa0e9093924.jpg	image	v1	2023-10-08 00:00:00	2023-10-11 12:34:31.438847
\.

ALTER TABLE ONLY public.apods
    ADD CONSTRAINT apods_date_key UNIQUE (date);

CREATE INDEX date_index ON public.apods USING btree (date);