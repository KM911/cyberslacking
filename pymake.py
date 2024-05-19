
import sys
import os
import subprocess

# Core function

# 输出cmd path or bash path

Github_Username = "KM911"


def Run(command: str):
    subprocess.run(command, shell=True)

# def modulelient(command: str):
#     # 重定向stdout


def ShowCommand(command: str):
    print(command)
    Run(command)


def CommandResult(command: str) -> str:
    return os.popen(command).read()


def RunSlient(command: str):

    subprocess.run(command, shell=True, stdout=subprocess.DEVNULL,
                   stderr=subprocess.DEVNULL)


def RunSlientVerbose(command: str):
    print(command)
    subprocess.run(command, shell=True, stdout=subprocess.DEVNULL,
                   stderr=subprocess.DEVNULL)


def RunSlientResult(command: str) -> str:
    file = open(os.path.join(os.environ.get("temp"), "temp.txt"), "w+")
    subprocess.run(command, shell=True, stdout=file,
                   stderr=subprocess.DEVNULL)
    file.seek(0)
    result = file.read()
    file.close()
    return result


def Cd(path: str):
    os.chdir(path)


def ExecutePathConvent(path: str) -> str:
    if os.name == "nt":
        return path
    else:
        return "./"+path


def ExeConvent(path: str) -> str:
    if os.name == "nt":
        return path+".exe"
    else:
        return path


def Load_Project_Env():
    global pwd, project, env
    pwd = os.getcwd()
    # print("pwd =", pwd)
    project = pwd.replace("\\", "/").split("/")[-1]
    # print("project =", project)
    env = os.environ.copy()


def Get_Programming_Using():
    global programming_using
    SupportLanguage = ["go", "rs", "c", "cpp", "py"]


def CheckEnv():
    print("need python3 rg")

# basic function


def help():  # show all function
    content = [x[4:-1] for x in open("pymake.py", "r", encoding='utf-8').readlines()
               if x.startswith("def") and x[4].islower()]
    print("\n".join(content))


def clean():
    # TODO  windows without rm
    Load_Project_Env()
    ShowCommand("rm -rf "+project)
    ShowCommand("rm -rf *.exe")
    ShowCommand("rm -rf *.test")
    ShowCommand("rm -rf *.out")
    ShowCommand("rm -rf *.log")
    ShowCommand("rm -rf *.prof")

# go function


def Load_Go_Env():
    Load_Project_Env()
    global mod
    mod = open("go.mod", "r").readline().split(" ")[
        1].strip()


def go_test():
    Load_Go_Env()
    Run(" ".join(["go", "test", "-v", mod+"/test"]))

# 一个大问题 它不会显示 关于内存占用的部分
# 131072 B/op 1 allocs/op


def Go_Benchmark(package: str):
    Load_Go_Env()
    Cd(package)
    command = " ".join(
        ["go", "test", "-c", "-o", "benchmark.test", mod+"/"+package])
    Run(command)
    command = " ".join([ExecutePathConvent("benchmark.test"),
                       "-test.benchmem", "-test.bench",  "."])

    result = CommandResult(command).split("\n")[4:-2]
    # do not need to show OS and CPU info
    # print("result =",result)
    # print("\n".join(result[:4]))
    # result = result[4:]

    import re
    regex_time = re.compile(r"\d+.?\d?(?= [mnu]s/op)")

    function_result = [result[i].split("-12") for i in range(len(result))]
    # 这里就需要进行判断
    # if function[0][1].find("B/op")!=0 :
    #     # min_memory = float(function_result[0][1].split(" ")[0])
    #     regex_memory = re.compile(r"\d+ B/op")
    #     regex_allocs = re.compile(r"\d+ allocs/op")
    #     function_result[i].append(regex_memory.findall(function_result[i][1])[0])

    # function_result[i].append(regex_allocs.findall(function_result[i][1])[0])

    unit = re.findall(r"[mnu]s/op", function_result[0][1],)[0]

    for i in range(len(function_result)):
        function_result[i][1] = float(
            regex_time.findall(function_result[i][1])[0])

    sorted_function_result = sorted(function_result, key=lambda x: float(x[1]))
    min_time = float(sorted_function_result[0][1])
    format_output = [
        f"{x[0]:26}    {x[1]:12}{unit}   {round(float(x[1]/min_time), 3):10}    {float(x[1])-min_time}{unit}" for x in sorted_function_result]

    print("\n".join(format_output))
    # os.remove("benchmark.test")


def go_bench():  # go test -bench
    global argv
    Go_Benchmark("benchmark")
    # if len(argv) == 0:
    #     Go_Benchmark('benchmark')
    # else:
    #     for i in argv:
    #         Go_Benchmark(i)
# TODO profile


def go_prof():
    # main or benchmark ????
    global argv
    if len(argv) == 0:
        pass
    # pass


def go_run():
    Load_Go_Env()
    Run(" ".join(["go", "run", "main.go"]+argv))


def go_build():
    Load_Project_Env()
    if os.name == "nt":
        go_win()
    else:
        go_linux()


def go_win():
    Load_Project_Env()
    env.update({"CGO_ENABLED": "0", "GOOS": "windows"})
    subprocess.run(["go", "build", "-ldflags", "-s -w",
                   "-o", project+".exe"], env=env)


def go_mod():
    # go: D:\GITHUB\KM911\template\p\gm\go.mod already exists
    Load_Project_Env()
    CompletedProcess = subprocess.run(["go", "mod", "init", "github.com/"+Github_Username +
                                      "/"+project], env=env, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    if CompletedProcess.returncode == 0:
        print("go mod init success")
    else:
        print("go mod init github.com/"+Github_Username+"/"+project)
        with open("go.mod", "r") as file:
            lines = file.readlines()
            lines[0] = "module github.com/"+Github_Username+"/"+project+"\n"
        with open("go.mod", "w") as file:
            file.writelines(lines)


def go_hidegui():
    Load_Project_Env()
    env.update({"CGO_ENABLED": "0", "GOOS": "windows"})
    subprocess.run(["go", "build", "-ldflags",
                   "-s -w -H=windowsgui", "-o", project+".exe"], env=env)


def go_linux():
    Load_Project_Env()
    env.update({"CGO_ENABLED": "0", "GOOS": "linux"})
    subprocess.run(["go", "build", "-ldflags",
                   "-s -w", "-o", project], env=env)


def go_static():
    Load_Project_Env()
    env.update({"CGO_ENABLED": "0", "GOOS": "linux"})
    subprocess.run(["go", "build", "-ldflags",
                   "-s -w -extldflags -static", "-o", project], env=env)


def go_proxy():
    ShowCommand("go env -w GOPROXY=https://goproxy.cn,direct")


def upx():
    ShowCommand("upx *.exe")


# Docker function

def Load_Docker_Env():
    Load_Project_Env()
    global image_info, version
    image_info = [x for x in CommandResult("docker image list").split(
        "\n") if x.startswith(project)]
    if len(image_info) == 0:
        print("no image")
        version = "0.0.0"
    else:
        import re
        version = re.findall(r"\d+\.\d+\.\d+", image_info[0])[0]


def image():
    Load_Project_Env()
    Load_Docker_Env()
    if len(image_info) == 0:
        print("no image")
    else:
        print(image_info)


def container():
    Load_Project_Env()
    container = [x for x in CommandResult(
        "docker container list").split("\n") if x.startswith(project)]
    if len(container) == 0:
        print("no container")
    else:
        print(container)


def docker_init():
    os.mkdir("docker")
    os.mkdir("docker/public")

    file = open("docker/Dockerfile", "w+")
    file.write("FROM alpine:latest\n")
    file.write("WORKDIR /app\n")
    file.write("ADD . /app\n")
    file.close()


def image_build():
    Load_Docker_Env()
    import re
    version_num = version.split(".")
    version_num[-1] = str(int(version_num[-1])+1)
    new_version = ".".join(version_num)

    if os.path.exists(project):
        if os.path.exists("docker/public/"+project):
            os.remove("docker/public/"+project)
        os.rename(project, "docker/public/"+project)

    Run("docker build -t "+project+":"+new_version+" ./docker/.")
    Run("docker image rm -f "+project+":"+version)


def container_clean():
    Run("docker container prune -f")


def image_clean():
    Run("docker image prune -f")


def image_run():
    Load_Docker_Env()
    Run("docker run -it "+project+":"+version)


# main function
# TODO add more typo
ErrorDict = {"iamge": "image", "benhc": "bench",
             "clnea": "clean", "dokcer": "docker", "dokecr": "docker"}

if __name__ == "__main__":
    global argv
    argc = len(sys.argv)
    if argc == 1:
        help()
    elif argc == 2:
        try:
            argv = sys.argv[2:]
            eval(sys.argv[1]+"()")
        except NameError:
            # fix typo
            fixed = ErrorDict.get(sys.argv[1], None)
            if fixed is None:
                help()
            else:
                eval(fixed+"()")
    else:
        argv = sys.argv[2:]
        try:
            eval("_".join(sys.argv[1:3])+"()")
        except NameError:
            fixed_1 = ErrorDict.get(sys.argv[1], None)
            fixed_2 = ErrorDict.get(sys.argv[2], None)
            if fixed_1 is None and fixed_2 is None:
                help()
            else:
                try:
                    if fixed_1 is None:
                        eval(sys.argv[1]+"_"+fixed_2+"()")
                    else:
                        eval(fixed_1+"_"+sys.argv[2]+"()")
                except NameError:
                    help()
