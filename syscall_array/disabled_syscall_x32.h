#ifndef  DISABLED_SYSCALL_INC
#define  DISABLED_SYSCALL_INC

const int disabled_syscall[318] = {
	//syscall number array comes from ZOJ
    0, /*0, restart_syscall*/
    0, /*1, exit*/
    1, /*2, fork*/
    0, /*3, read*/
    0, /*4, write*/
    0, /*5, open*/
    0, /*6, close*/
    0, /*7, waitpid*/
    1, /*8, creat*/
    1, /*9, link*/
    1, /*10, unlink*/
    1, /*11, execve*/
    1, /*12, chdir*/
    0, /*13, time*/
    1, /*14, mknod*/
    1, /*15, chmod*/
    0, /*16, lchown*/
    1, /*17, break*/
    1, /*18, oldstat*/
    0, /*19, lseek*/
    0, /*20, getpid*/
    1, /*21, mount*/
    1, /*22, umount*/
    0, /*23, setuid*/
    0, /*24, getuid*/
    0, /*25, stime*/
    1, /*26, ptrace*/
    0, /*27, alarm*/
    1, /*28, oldfstat*/
    1, /*29, pause*/
    0, /*30, utime*/
    1, /*31, stty*/
    1, /*32, gtty*/
    0, /*33, access*/
    0, /*34, nice*/
    1, /*35, ftime*/
    1, /*36, sync*/
    0, /*37, kill*/
    1, /*38, rename*/
    1, /*39, mkdir*/
    1, /*40, rmdir*/
    1, /*41, dup*/
    0, /*42, pipe*/
    0, /*43, times*/
    1, /*44, prof*/
    0, /*45, brk*/
    0, /*46, setgid*/
    0, /*47, getgid*/
    1, /*48, signal*/
    0, /*49, geteuid*/
    0, /*50, getegid*/
    0, /*51, acct*/
    1, /*52, umount2*/
    1, /*53, lock*/
    0, /*54, ioctl*/
    0, /*55, fcntl*/
    1, /*56, mpx*/
    0, /*57, setpgid*/
    1, /*58, ulimit*/
    1, /*59, oldolduname*/
    0, /*60, umask*/
    1, /*61, chroot*/
    0, /*62, ustat*/
    1, /*63, dup2*/
    0, /*64, getppid*/
    0, /*65, getpgrp*/
    0, /*66, setsid*/
    0, /*67, sigaction*/
    0, /*68, sgetmask*/
    0, /*69, ssetmask*/
    0, /*70, setreuid*/
    0, /*71, setregid*/
    0, /*72, sigsuspend*/
    0, /*73, sigpending*/
    0, /*74, sethostname*/
    1, /*75, setrlimit*/
    0, /*76, getrlimit*/
    0, /*77, getrusage*/
    0, /*78, gettimeofday*/
    0, /*79, settimeofday*/
    0, /*80, getgroups*/
    0, /*81, setgroups*/
    1, /*82, select*/
    1, /*83, symlink*/
    1, /*84, oldlstat*/
    0, /*85, readlink*/
    0, /*86, uselib*/
    1, /*87, swapon*/
    1, /*88, reboot*/
    1, /*89, readdir*/
    0, /*90, mmap*/
    0, /*91, munmap*/
    0, /*92, truncate*/
    0, /*93, ftruncate*/
    1, /*94, fchmod*/
    0, /*95, fchown*/
    0, /*96, getpriority*/
    0, /*97, setpriority*/
    1, /*98, profil*/
    1, /*99, statfs*/
    1, /*100, fstatfs*/
    0, /*101, ioperm*/
    0, /*102, socketcall*/
    1, /*103, syslog*/
    0, /*104, setitimer*/
    0, /*105, getitimer*/
    0, /*106, stat*/
    0, /*107, lstat*/
    0, /*108, fstat*/
    0, /*109, olduname*/
    0, /*110, iopl*/
    0, /*111, vhangup*/
    1, /*112, idle*/
    0, /*113, vm86old*/
    0, /*114, wait4*/
    1, /*115, swapoff*/
    0, /*116, sysinfo*/
    1, /*117, ipc*/
    0, /*118, fsync*/
    0, /*119, sigreturn*/
    1, /*120, clone*/
    0, /*121, setdomainname*/
    0, /*122, uname*/
    0, /*123, modify_ldt*/
    0, /*124, adjtimex*/
    0, /*125, mprotect*/
    0, /*126, sigprocmask*/
    1, /*127, create_module*/
    1, /*128, init_module*/
    1, /*129, delete_module*/
    0, /*130, get_kernel_syms*/
    0, /*131, quotactl*/
    0, /*132, getpgid*/
    0, /*133, fchdir*/
    0, /*134, bdflush*/
    0, /*135, sysfs*/
    0, /*136, personality*/
    1, /*137, afs_syscall*/
    0, /*138, setfsuid*/
    0, /*139, setfsgid*/
    0, /*140, llseek*/
    1, /*141, getdents*/
    1, /*142, newselect*/
    0, /*143, flock*/
    0, /*144, msync*/
    0, /*145, readv*/
    0, /*146, writev*/
    0, /*147, getsid*/
    0, /*148, fdatasync*/
    0, /*149, sysctl*/
    0, /*150, mlock*/
    0, /*151, munlock*/
    0, /*152, mlockall*/
    0, /*153, munlockall*/
    0, /*154, sched_setparam*/
    0, /*155, sched_getparam*/
    0, /*156, sched_setscheduler*/
    0, /*157, sched_getscheduler*/
    0, /*158, sched_yield*/
    0, /*159, sched_get_priority_max*/
    0, /*160, sched_get_priority_min*/
    0, /*161, sched_rr_get_interval*/
    1, /*162, nanosleep*/
    0, /*163, mremap*/
    0, /*164, setresuid*/
    0, /*165, getresuid*/
    0, /*166, vm86*/
    1, /*167, query_module*/
    0, /*168, poll*/
    0, /*169, nfsservctl*/
    0, /*170, setresgid*/
    0, /*171, getresgid*/
    0, /*172, prctl*/
    0, /*173, rt_sigreturn*/
    0, /*174, rt_sigaction*/
    0, /*175, rt_sigprocmask*/
    0, /*176, rt_sigpending*/
    0, /*177, rt_sigtimedwait*/
    0, /*178, rt_sigqueueinfo*/
    0, /*179, rt_sigsuspend*/
    1, /*180, pread64*/
    1, /*181, pwrite64*/
    1, /*182, chown*/
    0, /*183, getcwd*/
    0, /*184, capget*/
    0, /*185, capset*/
    1, /*186, sigaltstack*/
    1, /*187, sendfile*/
    1, /*188, getpmsg*/
    1, /*189, putpmsg*/
    1, /*190, vfork*/
    0, /*191, ugetrlimit*/
    0, /*192, mmap2*/
    0, /*193, truncate64*/
    0, /*194, ftruncate64*/
    0, /*195, stat64*/
    0, /*196, lstat64*/
    0, /*197, fstat64*/
    1, /*198, lchown32*/
    0, /*199, getuid32*/
    0, /*200, getgid32*/
    0, /*201, geteuid32*/
    0, /*202, getegid32*/
    0, /*203, setreuid32*/
    0, /*204, setregid32*/
    0, /*205, getgroups32*/
    0, /*206, setgroups32*/
    0, /*207, fchown32*/
    0, /*208, setresuid32*/
    0, /*209, getresuid32*/
    0, /*210, setresgid32*/
    0, /*211, getresgid32*/
    1, /*212, chown32*/
    0, /*213, setuid32*/
    0, /*214, setgid32*/
    0, /*215, setfsuid32*/
    0, /*216, setfsgid32*/
    0, /*217, pivot_root*/
    0, /*218, mincore*/
    0, /*219, madvise*/
    1, /*220, madvise1*/
    0, /*221, getdents64*/
    1, /*222, fcntl64*/
    1, /*223, unused*/
    0, /*224, gettid*/
    0, /*225, readahead*/
    1, /*226, setxattr*/
    1, /*227, lsetxattr*/
    1, /*228, fsetxattr*/
    1, /*229, getxattr*/
    1, /*230, lgetxattr*/
    1, /*231, fgetxattr*/
    1, /*232, listxattr*/
    1, /*233, llistxattr*/
    1, /*234, flistxattr*/
    1, /*235, removexattr*/
    1, /*236, lremovexattr*/
    1, /*237, fremovexattr*/
    0, /*238, tkill*/
    1, /*239, sendfile64*/
    1, /*240, futex*/
    1, /*241, sched_setaffinity*/
    1, /*242, sched_getaffinity*/
    0, /*243, set_thread_area*/
    0, /*244, get_thread_area*/
    1, /*245, io_setup*/
    1, /*246, io_destroy*/
    1, /*247, io_getevents*/
    1, /*248, io_submit*/
    1, /*249, io_cancel*/
    1, /*250, fadvise64*/
    1, /*251, unused*/
    0, /*252, exit_group*/
    1, /*253, lookup_dcookie*/
    1, /*254, epoll_create*/
    1, /*255, epoll_ctl*/
    1, /*256, epoll_wait*/
    1, /*257, remap_file_pages*/
    1, /*258, set_tid_address*/
    1, /*259, timer_create*/
    1, /*260, timer_settime*/
    1, /*261, timer_gettime*/
    1, /*262, timer_getoverrun*/
    1, /*263, timer_delete*/
    1, /*264, clock_settime*/
    1, /*265, clock_gettime*/
    1, /*266, clock_getres*/
    1, /*267, clock_nanosleep*/
    1, /*268, statfs64*/
    1, /*269, fstatfs64*/
    1, /*270, tgkill*/
    1, /*271, utimes*/
    1, /*272, fadvise64_64*/
    1, /*273, vserver*/
        1, /*274, mbind*/
        1, /*275, get_mempolicy*/
        1, /*276, set_mempolicy*/
        1, /*277, mq_open*/
        1, /*278, mq_unlink*/
        1, /*279, mq_timedsend*/
        1, /*280, mq_timedreceive*/
        1, /*281, mq_notify*/
        1, /*282, mq_getsetattr*/
        1, /*283, kexec_load*/
        1, /*284, waitid*/
        1, /*285, ni_syscall*/
        1, /*286, add_key*/
        1, /*287, request_key*/
        1, /*288, keyctl*/
        1, /*289, ioprio_set*/
        1, /*290, ioprio_get*/
        1, /*291, inotify_init*/
        1, /*292, inotify_add_watch*/
        1, /*293, inotify_rm_watch*/
        1, /*294, migrate_pages*/
        1, /*295, openat*/
        1, /*296, mkdirat*/
        1, /*297, mknodat*/
        1, /*298, fchownat*/
        1, /*299, futimesat*/
        1, /*300, fstatat64*/
        1, /*301, unlinkat*/
        1, /*302, renameat*/
        1, /*303, linkat*/
        1, /*304, symlinkat*/
        1, /*305, readlinkat*/
        1, /*306, fchmodat*/
        1, /*307, faccessat*/
        1, /*308, pselect6*/
        1, /*309, ppoll*/
        1, /*310, unshare*/
        1, /*311, set_robust_list*/
        1, /*312, get_robust_list*/
        1, /*313, splice*/
        1, /*314, sync_file_range*/
        1, /*315, tee*/
        1, /*316, vmsplice*/
        1, /*317, move_pages*/
};

const char *syscallName[351]={
	//syscall name below comes from CentOS7 unistd_32.h
"restart_syscall 0",
"exit 1",
"fork 2",
"read 3",
"write 4",
"open 5",
"close 6",
"waitpid 7",
"creat 8",
"link 9",
"unlink 10",
"execve 11",
"chdir 12",
"time 13",
"mknod 14",
"chmod 15",
"lchown 16",
"break 17",
"oldstat 18",
"lseek 19",
"getpid 20",
"mount 21",
"umount 22",
"setuid 23",
"getuid 24",
"stime 25",
"ptrace 26",
"alarm 27",
"oldfstat 28",
"pause 29",
"utime 30",
"stty 31",
"gtty 32",
"access 33",
"nice 34",
"ftime 35",
"sync 36",
"kill 37",
"rename 38",
"mkdir 39",
"rmdir 40",
"dup 41",
"pipe 42",
"times 43",
"prof 44",
"brk 45",
"setgid 46",
"getgid 47",
"signal 48",
"geteuid 49",
"getegid 50",
"acct 51",
"umount2 52",
"lock 53",
"ioctl 54",
"fcntl 55",
"mpx 56",
"setpgid 57",
"ulimit 58",
"oldolduname 59",
"umask 60",
"chroot 61",
"ustat 62",
"dup2 63",
"getppid 64",
"getpgrp 65",
"setsid 66",
"sigaction 67",
"sgetmask 68",
"ssetmask 69",
"setreuid 70",
"setregid 71",
"sigsuspend 72",
"sigpending 73",
"sethostname 74",
"setrlimit 75",
"getrlimit 76",
"getrusage 77",
"gettimeofday 78",
"settimeofday 79",
"getgroups 80",
"setgroups 81",
"select 82",
"symlink 83",
"oldlstat 84",
"readlink 85",
"uselib 86",
"swapon 87",
"reboot 88",
"readdir 89",
"mmap 90",
"munmap 91",
"truncate 92",
"ftruncate 93",
"fchmod 94",
"fchown 95",
"getpriority 96",
"setpriority 97",
"profil 98",
"statfs 99",
"fstatfs 100",
"ioperm 101",
"socketcall 102",
"syslog 103",
"setitimer 104",
"getitimer 105",
"stat 106",
"lstat 107",
"fstat 108",
"olduname 109",
"iopl 110",
"vhangup 111",
"idle 112",
"vm86old 113",
"wait4 114",
"swapoff 115",
"sysinfo 116",
"ipc 117",
"fsync 118",
"sigreturn 119",
"clone 120",
"setdomainname 121",
"uname 122",
"modify_ldt 123",
"adjtimex 124",
"mprotect 125",
"sigprocmask 126",
"create_module 127",
"init_module 128",
"delete_module 129",
"get_kernel_syms 130",
"quotactl 131",
"getpgid 132",
"fchdir 133",
"bdflush 134",
"sysfs 135",
"personality 136",
"afs_syscall 137",
"setfsuid 138",
"setfsgid 139",
"_llseek 140",
"getdents 141",
"_newselect 142",
"flock 143",
"msync 144",
"readv 145",
"writev 146",
"getsid 147",
"fdatasync 148",
"_sysctl 149",
"mlock 150",
"munlock 151",
"mlockall 152",
"munlockall 153",
"sched_setparam 154",
"sched_getparam 155",
"sched_setscheduler 156",
"sched_getscheduler 157",
"sched_yield 158",
"sched_get_priority_max 159",
"sched_get_priority_min 160",
"sched_rr_get_interval 161",
"nanosleep 162",
"mremap 163",
"setresuid 164",
"getresuid 165",
"vm86 166",
"query_module 167",
"poll 168",
"nfsservctl 169",
"setresgid 170",
"getresgid 171",
"prctl 172",
"rt_sigreturn 173",
"rt_sigaction 174",
"rt_sigprocmask 175",
"rt_sigpending 176",
"rt_sigtimedwait 177",
"rt_sigqueueinfo 178",
"rt_sigsuspend 179",
"pread64 180",
"pwrite64 181",
"chown 182",
"getcwd 183",
"capget 184",
"capset 185",
"sigaltstack 186",
"sendfile 187",
"getpmsg 188",
"putpmsg 189",
"vfork 190",
"ugetrlimit 191",
"mmap2 192",
"truncate64 193",
"ftruncate64 194",
"stat64 195",
"lstat64 196",
"fstat64 197",
"lchown32 198",
"getuid32 199",
"getgid32 200",
"geteuid32 201",
"getegid32 202",
"setreuid32 203",
"setregid32 204",
"getgroups32 205",
"setgroups32 206",
"fchown32 207",
"setresuid32 208",
"getresuid32 209",
"setresgid32 210",
"getresgid32 211",
"chown32 212",
"setuid32 213",
"setgid32 214",
"setfsuid32 215",
"setfsgid32 216",
"pivot_root 217",
"mincore 218",
"madvise 219",
"getdents64 220",
"fcntl64 221",
"gettid 224",
"readahead 225",
"setxattr 226",
"lsetxattr 227",
"fsetxattr 228",
"getxattr 229",
"lgetxattr 230",
"fgetxattr 231",
"listxattr 232",
"llistxattr 233",
"flistxattr 234",
"removexattr 235",
"lremovexattr 236",
"fremovexattr 237",
"tkill 238",
"sendfile64 239",
"futex 240",
"sched_setaffinity 241",
"sched_getaffinity 242",
"set_thread_area 243",
"get_thread_area 244",
"io_setup 245",
"io_destroy 246",
"io_getevents 247",
"io_submit 248",
"io_cancel 249",
"fadvise64 250",
"exit_group 252",
"lookup_dcookie 253",
"epoll_create 254",
"epoll_ctl 255",
"epoll_wait 256",
"remap_file_pages 257",
"set_tid_address 258",
"timer_create 259",
"timer_settime 260",
"timer_gettime 261",
"timer_getoverrun 262",
"timer_delete 263",
"clock_settime 264",
"clock_gettime 265",
"clock_getres 266",
"clock_nanosleep 267",
"statfs64 268",
"fstatfs64 269",
"tgkill 270",
"utimes 271",
"fadvise64_64 272",
"vserver 273",
"mbind 274",
"get_mempolicy 275",
"set_mempolicy 276",
"mq_open 277",
"mq_unlink 278",
"mq_timedsend 279",
"mq_timedreceive 280",
"mq_notify 281",
"mq_getsetattr 282",
"kexec_load 283",
"waitid 284",
"add_key 286",
"request_key 287",
"keyctl 288",
"ioprio_set 289",
"ioprio_get 290",
"inotify_init 291",
"inotify_add_watch 292",
"inotify_rm_watch 293",
"migrate_pages 294",
"openat 295",
"mkdirat 296",
"mknodat 297",
"fchownat 298",
"futimesat 299",
"fstatat64 300",
"unlinkat 301",
"renameat 302",
"linkat 303",
"symlinkat 304",
"readlinkat 305",
"fchmodat 306",
"faccessat 307",
"pselect6 308",
"ppoll 309",
"unshare 310",
"set_robust_list 311",
"get_robust_list 312",
"splice 313",
"sync_file_range 314",
"tee 315",
"vmsplice 316",
"move_pages 317",
"getcpu 318",
"epoll_pwait 319",
"utimensat 320",
"signalfd 321",
"timerfd_create 322",
"eventfd 323",
"fallocate 324",
"timerfd_settime 325",
"timerfd_gettime 326",
"signalfd4 327",
"eventfd2 328",
"epoll_create1 329",
"dup3 330",
"pipe2 331",
"inotify_init1 332",
"preadv 333",
"pwritev 334",
"rt_tgsigqueueinfo 335",
"perf_event_open 336",
"recvmmsg 337",
"fanotify_init 338",
"fanotify_mark 339",
"prlimit64 340",
"name_to_handle_at 341",
"open_by_handle_at 342",
"clock_adjtime 343",
"syncfs 344",
"sendmmsg 345",
"setns 346",
"process_vm_readv 347",
"process_vm_writev 348",
"kcmp 349",
"finit_module 350"
};
#endif   /* ----- #ifndef DISABLED_SYSCALL_INC  ----- */

