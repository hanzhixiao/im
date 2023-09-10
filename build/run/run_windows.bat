@echo off
cd %~p0../../
set service_names= auth chat chat_invite chat_member chat_msg msg_gateway server_mgr dist interfaces upload message msg_history cloud_msg user convo cache lbs
(for %%a in (%service_names%) do (
    start lark_%%a.exe
))