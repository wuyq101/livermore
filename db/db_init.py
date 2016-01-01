from yoyo import step

transaction(
    step(
        """
        create table if not exists stock(
            id bigint not null auto_increment,
            name varchar(64) not null comment '股票名称',
            code varchar(16) not null comment '股票代码',
            type varchar(10) not null default 'sh' comment '股票市场标识，sh sz',
            total_value bigint not null default 0 comment '总市值 单位分',
            circulate_value bigint not null default 0 comment '流通市值 单位分',
            primary key (id)
        ) ENGINE = InnoDB;
        """,
        """
        drop table if exists stock;
        """
    )
)
