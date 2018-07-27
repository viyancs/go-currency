This project is REST Full API for managing currency rates, To run these app, in root directory, run:

- mkdir go-currency
- cd go-currency
- git clone https://github.com/viyancs/go-currency.git .

$ export GOPATH=$PWD

Then building schema:

$ go build -o binschema schema.go

$ ./binschema

```bash
schema command usage:

        sample             - importing sample data
        orm syncdb -force  - syncdb model to table by droping table first
        orm syncdb         - syncdb model to table
        [options] -v       - verbose output
```

$ ./binschema orm syncdb -force


```bash
drop table `dailyexchange`
    DROP TABLE IF EXISTS `dailyexchange`

create table `dailyexchange`
    -- --------------------------------------------------
    --  Table Structure for `main.Dailyexchange`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `dailyexchange` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `c_from` varchar(3) NOT NULL DEFAULT '' ,
        `c_to` varchar(3) NOT NULL DEFAULT '' ,
        `rate` double precision  NULL DEFAULT 0 ,
        `date` datetime  NULL
    ) ENGINE=InnoDB;

```

$ ./binschema sample

```bash
Insert successfully = 6
```

$ go run app.go
