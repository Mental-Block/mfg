
#!/bin/bash
port=5432
server=172.18.0.2
database=mfg
user=root
password=root

for filename in ./*.sql; do
    [ -e "$filename" ] || continue
    sudo PGPASSWORD=$password psql -h $server -p $port -d $database -U $user < $filename
done

echo "finished"