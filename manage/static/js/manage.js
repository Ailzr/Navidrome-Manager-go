const apiUrl = 'music-list';
const token = localStorage.getItem('token');

// 获取用于显示音乐列表的元素
const musicListElement = document.getElementById('music-list');

// 定义一个函数来获取音乐列表并显示在前端页面上
function fetchMusicList() {
    // 创建一个新的XMLHttpRequest对象
    const xhr = new XMLHttpRequest();

    // 设置请求方法和URL
    xhr.open('GET', apiUrl);
    xhr.setRequestHeader('Authorization', token);

    // 设置响应类型为JSON
    xhr.responseType = 'json';

    // 当请求完成时执行的回调函数
    xhr.onload = function () {
        if (xhr.status === 200) {
            if (xhr.response.code === 200){
                const musicList = xhr.response.data;

                // 清空音乐列表元素的内容
                musicListElement.innerHTML = '';

                // 遍历音乐列表数据，将每个音乐项添加到音乐列表元素中
                musicList.forEach(function (music) {
                    const listItem = document.createElement('li');
                    listItem.textContent = music;
                    // 创建删除按钮
                    const deleteButton = document.createElement("button");
                    deleteButton.id = "deleteBtn";
                    deleteButton.textContent = "删除";
                    deleteButton.classList.add("deleteButton");

                    // 为删除按钮添加点击事件监听器
                    deleteButton.addEventListener("click", (event) => {
                        const li = event.target.parentElement
                        const songId = li.textContent.substring(0, li.textContent.length - 2)

                        const deleteUrl = "/manage/delete-music?music-name=" + encodeURIComponent(songId)
                        // 向后端发送删除请求
                        fetch(deleteUrl, {
                            method: 'DELETE',
                            headers: new Headers({
                                'Content-Type': 'application/json',
                                'Authorization': `${token}`
                            })
                        })
                            .then(response => response.json())
                            .then(data => {
                                if (data.code === 200){
                                    console.log('删除成功');
                                }else {
                                    console.log('删除失败' + data.message);
                                }
                            })
                            .catch((error) => {
                                console.error('请求错误：', error);
                            });

                        // 从列表中移除已删除的歌曲
                        li.remove();
                    });

                    // 将删除按钮添加到歌曲li元素后面
                    listItem.appendChild(deleteButton);
                    musicListElement.appendChild(listItem);
                });
            }
            else {
                console.error('请求失败，错误信息：', xhr.response.message);
            }
        } else {
            console.error('请求失败，状态码：', xhr.status);
        }
    };

    // 发送请求
    xhr.send();
}

// 调用fetchMusicList函数获取音乐列表并显示在前端页面上
fetchMusicList();

// scripts.js
document.getElementById('refresh-btn').addEventListener('click', function() {
    fetch(apiUrl, {
        method: 'GET',
        headers: new Headers({
            'Content-Type': 'application/json',
            'Authorization': `${token}`
        })
    })
        .then(response => response.json())
        .then(data => {
            const musicList = document.getElementById('music-list');
            musicList.innerHTML = ''; // 清空当前的音乐列表
            const music_list = data.data

            music_list.forEach(music => {
                const musicItem = document.createElement('li');
                musicItem.textContent = music;

                // 创建删除按钮
                const deleteButton = document.createElement("button");
                deleteButton.id = "deleteBtn";
                deleteButton.textContent = "删除";
                deleteButton.classList.add("deleteButton");

                // 为删除按钮添加点击事件监听器
                deleteButton.addEventListener("click", (event) => {
                    const li = event.target.parentElement
                    const songId = li.textContent.substring(0, li.textContent.length - 2)

                    const deleteUrl = "/manage/delete-music?music-name=" + encodeURIComponent(songId)
                    // 向后端发送删除请求
                    fetch(deleteUrl, {
                        method: 'DELETE',
                        headers: new Headers({
                            'Content-Type': 'application/json',
                            'Authorization': `${token}`
                        })
                    })
                        .then(response => response.json())
                        .then(data => {
                            if (data.code === 200){
                                console.log('删除成功');
                            }else {
                                console.log('删除失败' + data.message);
                            }
                        })
                        .catch((error) => {
                            console.error('请求错误：', error);
                        });
                    // 从列表中移除已删除的歌曲
                    li.remove();
                });

                // 将删除按钮添加到歌曲li元素后面
                musicList.appendChild(deleteButton);

                musicList.appendChild(musicItem);
            });
        })
        .catch(error => {
            console.error('Error fetching music list:', error);
        });
});


document.getElementById('clear').addEventListener('click', function() {
    const musicList = document.getElementById('music-list');
    musicList.innerHTML = '';
})

// const fileInput = document.getElementById('fileInput');
// const files = fileInput.files;
//
// for (let i = 0; i < files.length;i++){
//     const formData = new FormData();
//     const uploadUrl = "/manage/upload-music?file-name=" + files[i].name.substring(0, files[i].name.length - 4);
//     formData.append(files[i].name.substring(0, files[i].name.length - 4), files[i])
//     let xhr3 = new XMLHttpRequest();
//     xhr3.open('POST', uploadUrl, true);
//     xhr3.setRequestHeader("Authorization", `${token}`)
//     xhr3.upload.onloadstart = function() {
//         document.getElementById('status').innerHTML = "Uploading...";
//     }
//
//     xhr3.onreadystatechange = function (){
//         if (xhr3.readyState === XMLHttpRequest.DONE){
//             if (xhr3.status === 200){
//                 document.getElementById('status').innerText = 'Upload successful!';
//             } else {
//                 document.getElementById('status').innerText = 'Upload failed!';
//             }
//         }
//     }
//
//     xhr3.send(formData);
// }
//
// alert('上传结束');

function uploadFile(){

    const fileInput = document.getElementById('fileInput');
    const files = fileInput.files;
    const failList = [];
    let Fail = false;
    for (let i = 0; i < files.length;i++){
        const formData = new FormData();
        formData.append("file", files[i])
        let xhr3 = new XMLHttpRequest();
        xhr3.open('POST', "/manage/upload-music", true);
        xhr3.setRequestHeader("Authorization", `${token}`)

        xhr3.upload.onloadstart = function() {
            document.getElementById('status').className = 'uploading';
        }

        xhr3.onreadystatechange = function (){
            if (xhr3.readyState === XMLHttpRequest.DONE && xhr3.status === 200){
                let resp = JSON.parse(xhr3.responseText);
                if (resp.code !== 200){
                    Fail = true;
                    failList.push(files[i].name);
                }
            }
        }

        xhr3.send(formData);
    }

    if (!Fail){
        document.getElementById('status').className = 'uploadSuccess';
    }
    else {
        document.getElementById('status').className = 'uploadFail';
        const failSong = document.createElement('ul');
        for (let i = 0;i < failList.length;i++){
            const failItem = document.createElement('li');
            failItem.innerHTML = failList[i];
            failSong.appendChild(failItem);
        }
    }
}

