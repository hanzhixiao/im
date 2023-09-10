@echo off
set service_names= auth chat chat_invite chat_member chat_msg msg_gateway server_mgr dist interfaces upload message msg_history cloud_msg user convo cache lbs
(for %%a in (%service_names%) do (
    taskkill /f /im lark_%%a.exe
))