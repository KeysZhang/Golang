# Selpg

</br>

## Demand

#### 1.Install the  development environment of go language in Linux

#### 2. Complete the development foundation of CLI command in go language

</br>

## Instruction of Design

#### There are total five function to complete the tasks

####（1）`func get_args(...){...} `to get  the parameters and analize them from the command

#### （2）`func check_args(...){...}`to check the rationality of the parameters

#### （3）`func run(...){...}`to run the commands according the paramaters

#### （4）`func check_file_access(...){...}`to confirm the file existing and readable 

#### （5）`func check_err(...){...}`check all the errors and print them

#### （6）`func check_fout(...){...}`idetify the type of fout according the printDestination

#### （7）`func output_to_file(...){...}`to write the datas in file

####（8）`func  output_to_exc(...){...}`to call another process

</br>

## Test

###类型一

#### 1. $ selpg -s1 -e1 input_file

![1-1.png](/结果截图/1-1.png)

####2. $ selpg -s1 -e1 < input_file

![1-1.png](/结果截图/1-2.png)

#### 3. $ other_command | selpg -s10 -e20

![1-1.png](/结果截图/1-3.png)

####4. $ selpg -s10 -e20 input_file >output_file

![1-1.png](/结果截图/1-4-1.png)

![1-1.png](/结果截图/1-4-2.png)

#### 5. $ selpg -s10 -e20 input_file 2>error_file

![1-1.png](/结果截图/1-5-1.png)

![1-1.png](/结果截图/1-5-2.png)

#### 6. $ selpg -s10 -e20 input_file >output_file 2>error_file

![1-1.png](/结果截图/1-6-1.png)

![1-1.png](/结果截图/1-6-2.png)

![1-1.png](/结果截图/1-6-3.png)

![1-1.png](/结果截图/1-6-4.png)

#### 7. $ selpg -s10 -e20 input_file >output_file 2>/dev/null

![1-1.png](/结果截图/1-7.png)

#### 8. $ selpg -s10 -e20 input_file >/dev/null

![1-1.png](/结果截图/1-8.png)

#### 9. $ selpg -s10 -e20 input_file | other_command

![1-1.png](/结果截图/1-9.png)

#### 10. $ selpg -s10 -e20 input_file 2>error_file | other_command

![1-1.png](/结果截图/1-10.png)

### 类型二

####1. $ selpg -s10 -e20 -l66 input_file 

![1-1.png](/结果截图/2-1.png)

#### 2. $ selpg -s10 -e20 -f input_file

![1-1.png](/结果截图/2-2.png)

#### 3. $ selpg -s10 -e20 -dlp1 input_file

##### 	using the wf program to replace the lp. The wf program is used to write datas in wf_output.txt 

![1-1.png](/结果截图/3-1.png)

![1-1.png](/结果截图/3-2.png)

####4. $ selpg -s10 -e20 input_file > output_file 2>error_file &

![1-1.png](/结果截图/4.png)

















