const apiUploadFile = "http://localhost:8080/files/upload"
const btnFilePost = document.getElementById('uploadFile')
btnFilePost.addEventListener("click", async function(event){
    console.log("upload click")
    const node = document.getElementById("fileInputMultiPart");
    const data = new FormData()
    const file = node.files[0]

    data.append(file.name, file);
    fetch(apiUploadFile, {
        method: 'POST',
        body: data,
    }).then(resp => resp.text).then(console.log)
});