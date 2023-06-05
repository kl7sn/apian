# Brief Introduction

To solve the problem of service abnormal OOM, the OOM generated by sudden memory growth cann't be quickly located.
The purpose of this service is to help users automatically locate the OOM problem and provide the corresponding OOM analysis report.

# Module Introduction
- **Monitor Module**: Memory changes are detected every second
- **Forward Module**: The memory change data is sent to the server
  - Support OSS