#!/bin/bash

# Create the SQL file
cat > init.sql << EOL
-- Create the servers table
CREATE TABLE IF NOT EXISTS servers (
    id SERIAL PRIMARY KEY,
    model VARCHAR(255),
    ram VARCHAR(50),
    hdd VARCHAR(50),
    location VARCHAR(100),
    price VARCHAR(20)
);

-- Insert records
EOL

# Skip the header and process each line
tail -n +2 data.raw | while IFS=$'\t' read -r model ram hdd location price; do
    # Escape single quotes in the data
    model=$(echo "$model" | sed "s/'/''/g")
    ram=$(echo "$ram" | sed "s/'/''/g")
    hdd=$(echo "$hdd" | sed "s/'/''/g")
    location=$(echo "$location" | sed "s/'/''/g")
    price=$(echo "$price" | sed "s/'/''/g")
    
    # Add INSERT statement
    echo "INSERT INTO servers (model, ram, hdd, location, price) VALUES ('$model', '$ram', '$hdd', '$location', '$price');" >> init.sql
done

echo "SQL file generated successfully!"