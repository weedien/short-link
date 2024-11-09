import os

def batch_rename(dir_path, old_ext, new_ext):
    """
    批量重命名指定目录下的文件
    :param dir_path: 目录路径
    :param old_ext: 旧的文件扩展名(例如 ".txt")
    :param new_ext: 新的文件扩展名(例如 ".log")
    """
    try:
        # 获取目录下所有文件
        files = os.listdir(dir_path)

        # 遍历所有文件
        for file_name in files:
            # 如果文件扩展名为旧扩展名
            if file_name.endswith(old_ext):
                # 构建新文件名
                new_file_name = file_name.replace(old_ext, new_ext)

                # 构建旧文件完整路径
                old_file_path = os.path.join(dir_path, file_name)

                # 构建新文件完整路径
                new_file_path = os.path.join(dir_path, new_file_name)

                # 重命名文件
                os.rename(old_file_path, new_file_path)
                print(f"重命名 {file_name} 为 {new_file_name}")

        print("批量重命名成功!")
    except Exception as e:
        print(f"发生错误: {e}")

# 示例用法
batch_rename("../database/internal/common/persistence/po", ".gen.go", ".go")