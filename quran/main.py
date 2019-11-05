with open("../mysql/seeds/quran.sql", "w") as sql:
    with open("quran.csv", "r") as csv:
        csv.readline()
        for line in csv:
            line = line.strip()
            if (line == ""):
                continue
            items = line.split(",")
            if (len(items) < 4):
                continue
            sura_id = int(items[1])
            verse_id = int(items[2])
            ayah = str(items[3]).strip().strip('"')
            insert = 'INSERT IGNORE INTO quran (sura_id, verse_id, ayah) VALUES ({}, {}, "{}");'\
                .format(sura_id, verse_id, ayah)
            sql.write(insert + "\n")
