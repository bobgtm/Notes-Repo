# Unit 4

## Key Terms and Definitions

- Systemd
- Journalctl
- Cron / At / Systemd Timers
- Daemon
- Detection
- Response
- Mitigation
- Reporting
- Recovery
- Remediation
- Lessons Learned
- After action review
- Operations Bridge

## Notes during Lecture/Class: 

### Operating Systems
- Gathering system information
- Geneeral system info
- Disk that exist and filesystems
- System usage stats (iostat)
- Memory
- CPU
- Security
-- Logs
-- Open Ports (netstat -ntulp lsof -i :22)
-- processes (ps -ef | grep *service* | systemctl status | systemd-analyze blame )

Links:

Terms:

Useful tools:
- overleaf.com
- draw.io

## Discussion Posts

Discussion Post #1

Read this article: https://cio-wiki.org/wiki/Operations_Bridge

What terms and concepts are new to you?
- Event Correlation; Operations Bridge; Service Provisioning; 

Which pro seems the most important to you? Why?
- That an Operatons Bridge provides centralized management and streamlined operations seems the most important to me. Being able to consolidate infra management into a single location reigns in the sprawl of multi-service and infra management. Having to touch each service or piece of infrastructure across multiple platforms is time consuming and could lead to oversight if documentation and tracking is not put in place properly. 

Which con seems the most costly, or difficult to overcome to you? Why?

The con which seems most costly is the training and skill development required.
There seems to be a few routes a company can take depending on their needs but each
with their own inherent costs that make one option potentially more viable than the
others. Companies may want to hire another business to offer them services to help
with implementing or maintaining an Operations Bridge. Another approach might be to
facilitate training for team members already a part of the organization. Last, the
company may look to hire additional members to a team to bring them into the efforts
to integrate a OB. Each option could incur significant time and financial costs in
bringing people into the project who have the skills needed to use the OB
effectivly. If the skill set isn't there already, the time necessary to bring
everyone up to speed might be prohibitive for providing crucial or foundational
services to the company's clients.  


2. Your team has no documentation around how to check out a server during an incident. Write out a procedure of what you think an operations person should be doing on the system they suspect is not working properly. This may help, to get you started (https://zeltser.com/media/docs/security-incident-survey-cheat-sheet.pdf?msc=Cheat+Sheet+Blog) You may use AI for this, but let us know if you do.

1. Make note of the uptime
  1. Uptime or w

2. lastlog or last
3. arp -an 
3. Check system environment
    - cat /etc/\*release

## Digging Deeper

1. Read about battle drills here (https://en.wikipedia.org/wiki/Battle_drill)
Why might it be important to practice incident handling before an incident occurs?
Why might it be important to understand your tools before an incident occurs?


    

