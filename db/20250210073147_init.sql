-- +goose Up
-- +goose StatementBegin
CREATE TABLE activity (
                          id bigint GENERATED ALWAYS AS IDENTITY,
                          timespan_ms bigint,
                          timespan_sec bigint GENERATED ALWAYS AS (timespan_ms / 1000) STORED,
                          timespan_min numeric GENERATED ALWAYS AS (timespan_ms / 60000) STORED,
                          timespan_hour numeric GENERATED ALWAYS AS (timespan_ms / 3600000) STORED,
                          description text
);

-- Что ты мог бы сделать за 1 секунду
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (1000, 'Подмигнуть и улыбнуться'),
                                                    (1000, 'Крикнуть "БУ!" Застать кого-то врасплох.'),
                                                    (1000, 'Щелкнуть пальцами');

-- за 2 секунды
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (2000, 'Погладить кота (или себя, если кота нет).'),
                                                    (2000, 'Послать воздушный поцелуй. Просто потому что можно. 😘'),
                                                    (2000, 'Сделать странное лицо. Например, раздуть щеки или изобразить рыбу. 🐟');

-- за 3 секунды
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (3000, 'Сделать 18 прыжков через скакалку'),
                                                    (3000, 'Самое быстрое письмо "алфавита". Рекордсмены пишут от А до Я менее чем за 3 секунды. ✍️'),
                                                    (3000, 'Сказать что-то на трех языках. Например, "Привет!" — "Hello!" — "Hola!"');

-- за 5 секунд
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (5000, 'Сделать 5 глубоких вдохов. Полезно для перезагрузки.'),
                                                    (5000, 'Профессионалы, как Мэнни Пакьяо, могут нанести до 15-20 ударов за 5 секунд! 🥊'),
                                                    (5000, 'Самая быстрая жонглировка. Три мяча можно подбросить до 20 раз.');

-- за 1 минуту
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (60000, 'Сделать 10-15 отжиманий. Отличный способ взбодриться! 💪'),
                                                    (60000, 'Почистить зубы. И бонус — свежесть дыхания! 🪥'),
                                                    (60000, 'Количество поцелуев. В 2012 году пара поставила рекорд — 258 поцелуев за минуту. 💋');

-- за 2 минуты
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (120000, 'Скоростная речь. Один из рекордов — 655 слов за 120 секунд. 😲'),
                                                    (120000, 'Упаковка подарков. Мастера упаковывают до 10-12 подарков за 2 минуты! 🎁'),
                                                    (120000, 'Максимум отжиманий. Спортсмены делают до 150 отжиманий за 2 минуты. 💪');

-- за 10 минут
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (600000, 'Игра на инструменте. Освежите любимую мелодию. 🎸'),
                                                    (600000, 'Скоростное чтение. Прочитайте 5-10 страниц книги. 📖'),
                                                    (600000, 'Выпить воды и пройтись. Минимум 500 шагов вокруг комнаты или офиса. 🚶');

-- за 1 час
INSERT INTO activity (timespan_ms, description) VALUES
                                                    (3600000, 'Изучить новую тему. Посмотреть обучающее видео или прочитать несколько статей.'),
                                                    (3600000, 'Прочитать несколько глав книги. В среднем можно прочитать около 50-70 страниц. 📖'),
                                                    (3600000, 'Бег. За час профессиональные бегуны пробегают до 20 км.'),
                                                    (3600000, 'Спланировать поездку или отпуск.');


SELECT * from activity;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS activity;
-- +goose StatementEnd
