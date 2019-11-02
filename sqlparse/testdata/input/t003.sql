-- dialect: PostgreSQL
--- Numbers and Scientific notation
SELECT -2.134 * +5E+6 AS col1,
        78E-9 AS col2,
        .123E4 AS col3,
        ( 44.0 - 2 / 10 + 0.3 ) ^ 2  AS col4,
        -1 - 3 AS col5
    WHERE 1 = 2
        AND 4>3
        AND +5E+6<=7
 ;
