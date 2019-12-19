import glob

filenames = []
for filename in glob.glob("*.csv"):
    filename = filename.split(".")[0]
    filenames.append(filename.split(".")[0])
for filename in filenames:
    with open(f"../mysql/seeds/s002-{filename}.sql", "w") as sql:
        with open(f"{filename}.csv", "r") as csv:
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
                translation = ','.join(items[3:]).strip().strip('"')
                insert = 'INSERT IGNORE INTO translation (sura_id, verse_id, name, translation) VALUES ({}, {}, "{}", "{}");'\
                    .format(sura_id, verse_id, filename, translation)
                sql.write(insert + "\n")
