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
            unique key(code,market_day)
        ) ENGINE = InnoDB;
        """,
        """
        drop table if exists stock_daily;
        """
    ),
    step(
        """
        create table if not exists money_flow(
            id bigint not null auto_increment,
            name varchar(64) not null,
            code varchar(16) not null,
            closing_price bigint not null default 0,
            increase_rate bigint not null default 0,
            turnover_rate bigint not null default 0 comment '换手率',
            net_amount bigint not null default 0 comment '净流入，单位分',
            ratio_amount bigint not null default 0 comment '净流入 比率，1587 --> 15.87%',
            r0 bigint not null default 0 comment '超大单成交额',
            r1 bigint not null default 0 comment '大单成交额',
            r2 bigint not null default 0 comment '小单成交额',
            r3 bigint not null default 0 comment '散单成交额',
            r0_net bigint not null default 0 comment '超大单净流入',
            r1_net bigint not null default 0 comment '大单净流入',
            r2_net bigint not null default 0 comment '小单净流入',
            r3_net bigint not null default 0 comment '散单净流入',
            market_day varchar(8) not null comment '时间，格式20060102',
            created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            primary key(id),
            unique key(code,market_day)
        ) ENGINE = InnoDB;
        """,
        """
        drop table if exists money_flow;    
        """
    )
)
