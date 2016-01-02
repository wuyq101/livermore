from yoyo import step

transaction(
    step(
        """
        create table if not exists stock(
            id bigint not null auto_increment,
            name varchar(64) not null comment '股票名称',
            code varchar(16) not null comment '股票代码',
            total_capital bigint not null default 0 comment '总股本 单位股',
            curr_capital bigint not null default 0 comment '流通股 单位股',
            created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            primary key (id),
            unique key (code),
            unique key (name)
        ) ENGINE = InnoDB;
        """,
        """
        drop table if exists stock;
        """
    ),
    step(
        """
        create table if not exists stock_daily(
            id bigint not null auto_increment,
            name varchar(64) not null,
            code varchar(16) not null,
            cur_price bigint not null default 0 comment '当前价格，单位分',
            increase_amount bigint not null default 0 comment '涨跌额',
            increase_rate bigint not null default 0 comment '涨跌幅',
            high_price bigint not null default 0,
            low_price bigint not null default 0,
            turnover bigint not null default 0 comment '成交额',
            volumn bigint not null default 0 comment '成交量 手',
            main_buy bigint not null default 0 comment '主力买入',
            main_buy_rate bigint not null default 0 comment '主力买入占比',
            main_sell bigint not null default 0 comment '主力卖出',
            main_sell_rate bigint not null default 0 comment '主力卖出占比',
            individual_buy bigint not null default 0 comment '散户买入',
            individual_buy_rate bigint not null default 0,
            individual_sell bigint not null default 0,
            individual_sell_rate bigint not null default 0,
            market_day  varchar(8) not null comment '时间，格式20060102',
            created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            primary key(id),
            unique key(name,market_day)
        ) ENGINE = InnoDB;
        """,
        """
        drop table if exists stock_daily;
        """
    )
)
