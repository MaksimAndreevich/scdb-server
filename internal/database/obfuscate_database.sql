-- ============================================================================
-- ПОЛНАЯ ОБФУСКАЦИЯ ТАБЛИЦЫ education_organizations
-- ============================================================================
-- ВНИМАНИЕ: Этот скрипт ИЗМЕНИТ данные в таблице education_organizations!
-- ============================================================================

BEGIN;

-- Создаем временные функции для обфускации
-- ============================================================================

-- Обфускация телефона (скрываем последние 6 символов)
CREATE OR REPLACE FUNCTION temp_obfuscate_phone(original_phone TEXT)
RETURNS TEXT AS $$
BEGIN
    IF original_phone IS NULL OR original_phone = '' THEN
        RETURN original_phone;
    END IF;
    
    -- Если телефон короче 6 символов, возвращаем как есть
    IF LENGTH(original_phone) <= 6 THEN
        RETURN original_phone;
    END IF;
    
    -- Скрываем последние 6 символов
    RETURN LEFT(original_phone, LENGTH(original_phone) - 6) || 'XXXXXX';
END;
$$ LANGUAGE plpgsql;

-- Обфускация имени руководителя (скрываем половину текста)
CREATE OR REPLACE FUNCTION temp_obfuscate_head_name(original_name TEXT)
RETURNS TEXT AS $$
DECLARE
    name_length INTEGER;
    visible_part TEXT;
BEGIN
    IF original_name IS NULL OR original_name = '' THEN
        RETURN original_name;
    END IF;
    
    name_length := LENGTH(original_name);
    
    -- Если имя короче 4 символов, возвращаем как есть
    IF name_length <= 4 THEN
        RETURN original_name;
    END IF;
    
    -- Показываем первую половину, остальное скрываем
    visible_part := LEFT(original_name, name_length / 2);
    
    RETURN visible_part || REPEAT('*', name_length - (name_length / 2));
END;
$$ LANGUAGE plpgsql;

-- Обфускация ИНН/КПП/ОГРН (скрываем все кроме первых 2 символов)
CREATE OR REPLACE FUNCTION temp_obfuscate_legal_number(original_number TEXT)
RETURNS TEXT AS $$
BEGIN
    IF original_number IS NULL OR original_number = '' THEN
        RETURN original_number;
    END IF;
    
    -- Если номер короче 3 символов, возвращаем как есть
    IF LENGTH(original_number) <= 2 THEN
        RETURN original_number;
    END IF;
    
    -- Показываем первые 2 символа, остальное скрываем
    RETURN LEFT(original_number, 2) || REPEAT('X', LENGTH(original_number) - 2);
END;
$$ LANGUAGE plpgsql;

-- Обфускация short_name (скрываем половину текста)
CREATE OR REPLACE FUNCTION temp_obfuscate_short_name(original_name TEXT)
RETURNS TEXT AS $$
DECLARE
    name_length INTEGER;
    visible_part TEXT;
BEGIN
    IF original_name IS NULL OR original_name = '' THEN
        RETURN original_name;
    END IF;
    
    name_length := LENGTH(original_name);
    
    -- Если название короче 4 символов, возвращаем как есть
    IF name_length <= 4 THEN
        RETURN original_name;
    END IF;
    
    -- Показываем первую половину, остальное скрываем
    visible_part := LEFT(original_name, name_length / 2);
    
    RETURN visible_part || REPEAT('*', name_length - (name_length / 2));
END;
$$ LANGUAGE plpgsql;

-- Обфускация адреса (скрываем последние 5 слов)
CREATE OR REPLACE FUNCTION temp_obfuscate_address(original_address TEXT)
RETURNS TEXT AS $$
DECLARE
    words TEXT[];
    word_count INTEGER;
    visible_words INTEGER;
    result TEXT;
BEGIN
    IF original_address IS NULL OR original_address = '' THEN
        RETURN original_address;
    END IF;
    
    -- Разбиваем адрес на слова по запятой
    words := string_to_array(original_address, ',');
    word_count := array_length(words, 1);
    
    -- Если слов меньше или равно 5, скрываем все кроме первого
    IF word_count <= 5 THEN
        IF word_count = 1 THEN
            RETURN words[1];
        ELSE
            RETURN words[1] || ',XXXXXX';
        END IF;
    END IF;
    
    -- Показываем первые (общее_количество - 5) слов
    visible_words := word_count - 5;
    
    -- Собираем видимую часть адреса
    result := array_to_string(words[1:visible_words], ',');
    
    -- Добавляем скрытую часть
    RETURN result || ',XXXXXX';
END;
$$ LANGUAGE plpgsql;

-- ============================================================================
-- ОБФУСКАЦИЯ ДАННЫХ В ТАБЛИЦЕ education_organizations
-- ============================================================================

UPDATE education_organizations SET
    -- 1. Скрываем последние 5 слов в адресе
    post_address = temp_obfuscate_address(post_address),
    
    -- 2. Скрываем последние 6 символов телефона
    phone = temp_obfuscate_phone(phone),
    
    -- 3. Полностью скрываем факс
    fax = CASE 
        WHEN fax IS NOT NULL AND fax != '' THEN 'СКРЫТО'
        ELSE fax 
    END,
    
    -- 4. Полностью скрываем email
    email = CASE 
        WHEN email IS NOT NULL AND email != '' THEN 'СКРЫТО'
        ELSE email 
    END,
    
    -- 5. Полностью скрываем web_site
    web_site = CASE 
        WHEN web_site IS NOT NULL AND web_site != '' THEN 'СКРЫТО'
        ELSE web_site 
    END,
    
    -- 6. Скрываем половину текста в head_name
    head_name = temp_obfuscate_head_name(head_name),
    
    -- 7. Скрываем ОГРН (первые 2 символа + X)
    ogrn = temp_obfuscate_legal_number(ogrn),
    
    -- 8. Скрываем ИНН (первые 2 символа + X)
    inn = temp_obfuscate_legal_number(inn),
    
    -- 9. Скрываем КПП (первые 2 символа + X)
    kpp = temp_obfuscate_legal_number(kpp),
    
    -- 10. Скрываем половину short_name
    short_name = temp_obfuscate_short_name(short_name);

-- ============================================================================
-- ОЧИСТКА ВРЕМЕННЫХ ФУНКЦИЙ
-- ============================================================================

DROP FUNCTION IF EXISTS temp_obfuscate_phone(TEXT);
DROP FUNCTION IF EXISTS temp_obfuscate_head_name(TEXT);
DROP FUNCTION IF EXISTS temp_obfuscate_legal_number(TEXT);
DROP FUNCTION IF EXISTS temp_obfuscate_short_name(TEXT);
DROP FUNCTION IF EXISTS temp_obfuscate_address(TEXT);

-- ============================================================================
-- СТАТИСТИКА ПОСЛЕ ОБФУСКАЦИИ
-- ============================================================================

SELECT 
    'СТАТИСТИКА ОБФУСКАЦИИ:' as info,
    COUNT(*) as total_organizations,
    COUNT(CASE WHEN post_address LIKE '%XXXXXX' THEN 1 END) as addresses_obfuscated,
    COUNT(CASE WHEN phone LIKE '%XXXXXX' THEN 1 END) as phones_obfuscated,
    COUNT(CASE WHEN fax = 'СКРЫТО' THEN 1 END) as faxes_hidden,
    COUNT(CASE WHEN email = 'СКРЫТО' THEN 1 END) as emails_hidden,
    COUNT(CASE WHEN web_site = 'СКРЫТО' THEN 1 END) as websites_hidden,
    COUNT(CASE WHEN head_name LIKE '%*%' THEN 1 END) as head_names_obfuscated,
    COUNT(CASE WHEN ogrn LIKE '%XX%' THEN 1 END) as ogrn_obfuscated,
    COUNT(CASE WHEN inn LIKE '%XX%' THEN 1 END) as inn_obfuscated,
    COUNT(CASE WHEN kpp LIKE '%XX%' THEN 1 END) as kpp_obfuscated,
    COUNT(CASE WHEN short_name LIKE '%*%' THEN 1 END) as short_names_obfuscated
FROM education_organizations;

COMMIT;

-- ============================================================================
-- ПРОВЕРОЧНЫЕ ЗАПРОСЫ (выполните после основного скрипта)
-- ============================================================================

-- Проверка результатов обфускации
SELECT 'Примеры обфускированных данных:' as check_title;
SELECT 
    full_name,
    short_name,
    post_address,
    phone,
    fax,
    email,
    web_site,
    head_name,
    ogrn,
    inn,
    kpp
FROM education_organizations 
LIMIT 10;