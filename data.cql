insert into collection_config (sysid,id,name,interval,latency,metric,host,qry,updated_by,tsp)
values (uuid(),'yockgen-0001','power_consumption',10,10,'powerstat_package_current_power_consumption_watts','192.168.1.107','','yockgen',toTimestamp(now())); 

insert into collection_config (sysid,id,name,interval,latency,metric,host,qry,updated_by,tsp)
values (uuid(),'yockgen-0001','power_consumption',10,10,'cpu_usage_idle','192.168.1.107','','yockgen',toTimestamp(now())); 

insert into collection_config (sysid,id,name,interval,latency,metric,host,qry,updated_by,tsp)
values (uuid(),'samuel-0001','cpu_consumption',10,10,'cpu_usage_idle','192.168.1.108','','samtaripin',toTimestamp(now())); 

insert into collection_config (sysid,id,name,interval,latency,metric,host,qry,updated_by,tsp)
values (uuid(),'yockgen-0002','common_metrices',10,10,'cpu_usage_idle','TGL01','','yockgen',toTimestamp(now())); 

insert into collection_config (sysid,id,name,interval,latency,metric,host,qry,updated_by,tsp)
values (uuid(),'yockgen-0002','common_metrices',10,10,'mem_used_percent','TGL01','','yockgen',toTimestamp(now())); 

insert into collection_config (sysid,id,name,interval,latency,metric,host,qry,updated_by,tsp)
values (uuid(),'yockgen-0002','common_metrices',10,10,'cpu_usage_idle','KBL01','','yockgen',toTimestamp(now())); 

insert into collection_config (sysid,id,name,interval,latency,metric,host,qry,updated_by,tsp)
values (uuid(),'yockgen-0002','common_metrices',10,10,'mem_used_percent','KBL01','','yockgen',toTimestamp(now())); 
