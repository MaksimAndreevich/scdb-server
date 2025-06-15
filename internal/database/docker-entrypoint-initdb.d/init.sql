-- Федеральные округа 
CREATE TABLE IF NOT EXISTS federal_districts (
    id SERIAL PRIMARY KEY,
    short_name VARCHAR(10) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Регионы 
CREATE TABLE IF NOT EXISTS regions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    label VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL,
    type_short VARCHAR(10) NOT NULL,
    content_type VARCHAR(50) NOT NULL,
    kladr_id VARCHAR(20) NOT NULL,
    okato VARCHAR(20) NOT NULL,
    oktmo VARCHAR(20) NOT NULL,
    guid UUID NOT NULL,
    code VARCHAR(10) NOT NULL,
    iso_3166_2 VARCHAR(10) NOT NULL,
    population INTEGER,
    year_founded INTEGER,
    area DECIMAL(12,2),
    fullname VARCHAR(200) NOT NULL,
    name_en VARCHAR(200) NOT NULL,
    fk_federal_district_id INTEGER NOT NULL REFERENCES federal_districts(id),
    
    -- Падежи названия
    namecase_nominative VARCHAR(200) NOT NULL,
    namecase_genitive VARCHAR(200) NOT NULL,
    namecase_dative VARCHAR(200) NOT NULL,
    namecase_accusative VARCHAR(200) NOT NULL,
    namecase_ablative VARCHAR(200) NOT NULL,
    namecase_prepositional VARCHAR(200) NOT NULL,
    namecase_locative VARCHAR(200) NOT NULL,
    
    -- Столица региона
    capital_name VARCHAR(100) NOT NULL,
    capital_label VARCHAR(50) NOT NULL,
    capital_kladr_id VARCHAR(20) NOT NULL,
    capital_okato VARCHAR(20) NOT NULL,
    capital_oktmo VARCHAR(20) NOT NULL,
    capital_content_type VARCHAR(50) NOT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT regions_kladr_id_key UNIQUE (kladr_id), 
    CONSTRAINT regions_okato_key UNIQUE (okato),
    CONSTRAINT regions_oktmo_key UNIQUE (oktmo),
    CONSTRAINT regions_guid_key UNIQUE (guid),
    CONSTRAINT regions_code_key UNIQUE (code)
);

-- Города 
CREATE TABLE IF NOT EXISTS cities (
    fias_id TEXT PRIMARY KEY,
    address TEXT,
    postal_code TEXT,
    country TEXT,
    federal_district_name TEXT,
    fk_region_id INTEGER NOT NULL REFERENCES regions(id),
    fk_federal_district_id INTEGER NOT NULL REFERENCES federal_districts(id),
    region_type TEXT,
    region TEXT,
    area_type TEXT,
    area TEXT,
    city_type TEXT,
    city TEXT,
    kladr_id TEXT,
    fias_level INTEGER,
    capital_marker INTEGER,
    okato TEXT,
    oktmo TEXT,
    tax_office TEXT,
    timezone TEXT,
    geo_lat DOUBLE PRECISION,
    geo_lon DOUBLE PRECISION,
    population INTEGER,
    foundation_year INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Типы образовательных организицаций 
CREATE TABLE education_types (
    key TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    level TEXT NOT NULL,
    ownership_forms TEXT[] NOT NULL,
    keywords TEXT[] NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Образовательные организации 
CREATE TABLE IF NOT EXISTS education_organizations (
    id TEXT PRIMARY KEY,
    full_name TEXT,
    short_name TEXT,
    head_edu_org_id TEXT,
    is_branch BOOLEAN,
    post_address TEXT NOT NULL, 
    phone TEXT,
    fax TEXT,
    email TEXT,
    web_site TEXT,
    ogrn TEXT,
    inn TEXT,
    kpp TEXT,
    head_post TEXT,
    head_name TEXT,
    form_name TEXT,
    kind_name TEXT,
    type_name TEXT,
    fk_city_id TEXT REFERENCES cities(fias_id),
    fk_region_id INTEGER REFERENCES regions(id),
    fk_federal_district_id INTEGER REFERENCES federal_districts(id),
    fk_education_type_key TEXT REFERENCES education_types(key),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

