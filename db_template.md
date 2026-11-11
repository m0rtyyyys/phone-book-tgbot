phonebook-# \dt
               Список таблиц
 Схема  |     Имя     |   Тип   | Владелец 
--------+-------------+---------+----------
 public | cities      | таблица | postgres
 public | main        | таблица | postgres
 public | names       | таблица | postgres
 public | patronymics | таблица | postgres
 public | streets     | таблица | postgres
 public | surnames    | таблица | postgres
(6 строк)

phonebook-# \d+ main
                                                                           Таблица "public.main"
  Столбец   |          Тип          | Правило сортировки | Допустимость NULL |             По умолчанию              | Хранилище | Сжатие | Цель для статистики | Описани
е 
------------+-----------------------+--------------------+-------------------+---------------------------------------+-----------+--------+---------------------+--------
--
 uniq_id    | integer               |                    | not null          | nextval('main_uniq_id_seq'::regclass) | plain     |        |                     | 
 surname    | integer               |                    |                   |                                       | plain     |        |                     | 
 name_      | integer               |                    |                   |                                       | plain     |        |                     | 
 patronymic | integer               |                    |                   |                                       | plain     |        |                     | 
 street     | integer               |                    |                   |                                       | plain     |        |                     | 
 house_num  | character varying(5)  |                    |                   |                                       | extended  |        |                     | 
 bldg       | character varying(5)  |                    |                   |                                       | extended  |        |                     | 
 flat       | integer               |                    |                   |                                       | plain     |        |                     | 
 tel_num    | character varying(20) |                    |                   |                                       | extended  |        |                     | 
 city       | integer               |                    |                   |                                       | plain     |        |                     | 
Индексы:
    "main_pkey" PRIMARY KEY, btree (uniq_id)
Ограничения внешнего ключа:
    "main_city_fkey" FOREIGN KEY (city) REFERENCES cities(c_id)
    "main_name__fkey" FOREIGN KEY (name_) REFERENCES names(n_id)
    "main_patronymic_fkey" FOREIGN KEY (patronymic) REFERENCES patronymics(p_id)
    "main_street_fkey" FOREIGN KEY (street) REFERENCES streets(str_id)
    "main_surname_fkey" FOREIGN KEY (surname) REFERENCES surnames(s_id)
Ограничения NOT NULL:
    "main_uniq_id_not_null" NOT NULL "uniq_id"
Метод доступа: heap

phonebook-# \d+ cities
                                                                        Таблица "public.cities"
 Столбец |          Тип          | Правило сортировки | Допустимость NULL |             По умолчанию             | Хранилище | Сжатие | Цель для статистики | Описание 
---------+-----------------------+--------------------+-------------------+--------------------------------------+-----------+--------+---------------------+----------
 c_id    | integer               |                    | not null          | nextval('cities_c_id_seq'::regclass) | plain     |        |                     | 
 city    | character varying(25) |                    |                   |                                      | extended  |        |                     | 
Индексы:
    "cities_pkey" PRIMARY KEY, btree (c_id)
Ссылки извне:
    TABLE "main" CONSTRAINT "main_city_fkey" FOREIGN KEY (city) REFERENCES cities(c_id)
Ограничения NOT NULL:
    "cities_c_id_not_null" NOT NULL "c_id"
Метод доступа: heap

phonebook-# \d+ names
                                                                        Таблица "public.names"
 Столбец |          Тип          | Правило сортировки | Допустимость NULL |            По умолчанию             | Хранилище | Сжатие | Цель для статистики | Описание 
---------+-----------------------+--------------------+-------------------+-------------------------------------+-----------+--------+---------------------+----------
 n_id    | integer               |                    | not null          | nextval('names_n_id_seq'::regclass) | plain     |        |                     | 
 name    | character varying(25) |                    |                   |                                     | extended  |        |                     | 
Индексы:
    "names_pkey" PRIMARY KEY, btree (n_id)
Ссылки извне:
    TABLE "main" CONSTRAINT "main_name__fkey" FOREIGN KEY (name_) REFERENCES names(n_id)
Ограничения NOT NULL:
    "names_n_id_not_null" NOT NULL "n_id"
Метод доступа: heap

phonebook-# \d+ patronymics
                                                                         Таблица "public.patronymics"
  Столбец   |          Тип          | Правило сортировки | Допустимость NULL |               По умолчанию                | Хранилище | Сжатие | Цель для статистики | Опи
сание 
------------+-----------------------+--------------------+-------------------+-------------------------------------------+-----------+--------+---------------------+----
------
 p_id       | integer               |                    | not null          | nextval('patronymics_p_id_seq'::regclass) | plain     |        |                     | 
 patronymic | character varying(25) |                    |                   |                                           | extended  |        |                     | 
Индексы:
    "patronymics_pkey" PRIMARY KEY, btree (p_id)
Ссылки извне:
    TABLE "main" CONSTRAINT "main_patronymic_fkey" FOREIGN KEY (patronymic) REFERENCES patronymics(p_id)
Ограничения NOT NULL:
    "patronymics_p_id_not_null" NOT NULL "p_id"
Метод доступа: heap

phonebook-# \d+ surnames
                                                                        Таблица "public.surnames"
 Столбец |          Тип          | Правило сортировки | Допустимость NULL |              По умолчанию              | Хранилище | Сжатие | Цель для статистики | Описание 
---------+-----------------------+--------------------+-------------------+----------------------------------------+-----------+--------+---------------------+----------
 s_id    | integer               |                    | not null          | nextval('surnames_s_id_seq'::regclass) | plain     |        |                     | 
 surname | character varying(25) |                    |                   |                                        | extended  |        |                     | 
Индексы:
    "surnames_pkey" PRIMARY KEY, btree (s_id)
Ссылки извне:
    TABLE "main" CONSTRAINT "main_surname_fkey" FOREIGN KEY (surname) REFERENCES surnames(s_id)
Ограничения NOT NULL:
    "surnames_s_id_not_null" NOT NULL "s_id"
Метод доступа: heap

phonebook-# \d+ streets
                                                                         Таблица "public.streets"
 Столбец |          Тип          | Правило сортировки | Допустимость NULL |              По умолчанию               | Хранилище | Сжатие | Цель для статистики | Описание
 
---------+-----------------------+--------------------+-------------------+-----------------------------------------+-----------+--------+---------------------+---------
-
 str_id  | integer               |                    | not null          | nextval('streets_str_id_seq'::regclass) | plain     |        |                     | 
 street  | character varying(25) |                    |                   |                                         | extended  |        |                     | 
Индексы:
    "streets_pkey" PRIMARY KEY, btree (str_id)
Ссылки извне:
    TABLE "main" CONSTRAINT "main_street_fkey" FOREIGN KEY (street) REFERENCES streets(str_id)
Ограничения NOT NULL:
    "streets_str_id_not_null" NOT NULL "str_id"
Метод доступа: heap
