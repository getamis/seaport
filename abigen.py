import os

contract_folder = "contracts"
build_folder = "build"
output_folder = "output"
ignore_array = ["Mock", "Reference"]

def makedirs(outFolder):
    if not os.path.exists(outFolder):
        os.makedirs(outFolder)

def do_abigen():
    for root, dirs, files in os.walk(contract_folder):
        for f in files:
            exec = True
            if ".sol" not in f:
                continue
            for ignore in ignore_array:
                if ignore in f:
                    exec = False
            if exec:
                target = os.path.join(root, f)
                cmd = f'solc --allow-paths . --optimize @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ @rari-capital/=$(pwd)/node_modules/@rari-capital/ --abi {target} --bin {target} -o {build_folder} --overwrite'
                print(cmd)
                os.system(cmd)
                packageName = f.replace(".sol", "")
                abiFile =  os.path.join(build_folder, f.replace(".sol",".abi"))
                binFile =  os.path.join(build_folder, f.replace(".sol",".bin"))
                file_folder = os.path.join(output_folder, packageName)
                if not os.path.exists(file_folder):
                    os.makedirs(file_folder)
                goFile = os.path.join(file_folder, f.replace(".sol",".go"))
                cmd = f'abigen --abi={abiFile} --bin={binFile} --pkg={packageName} --out={goFile}'
                print(cmd)
                os.system(cmd)

if __name__ == "__main__":
    do_abigen()