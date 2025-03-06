
#!/bin/bash
port=5432
server=172.18.0.2
database=mfg
user=root
password=root

echo "Are you sure you want to delete all data? type YES to confirm"

read AreYouSure

if [[ "$AreYouSure" == "YES" ]]
then
for filename in ./*.sql; do
    [ -e "$filename" ] || continue

    table() {
        echo ${filename:4} | cut -f 1 -d '.'
    }

    sudo PGPASSWORD=$password psql -h $server -p $port -d $database -U $user -c "TRUNCATE TABLE public.$(table) CASCADE;"
done
fi

echo "finished"

    
