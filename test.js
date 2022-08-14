function trick(){
    let i, isum = 0
    for(i = 0; i < 1e9; i++){
        isum += i
    }
    // console.log(isum)
    console.log("successQ!!")
    return isum    
}


async function longProcess(){
    let i, isum = 0
    for(i = 0; i < 1e9; i++){
        isum += i
    }
    // console.log(isum)
    console.log("successQ!!")
    return isum   
    // return trick

    // return new Promise((resolve, reject) =>  {
    //     if(true){
    //         resolve(() => {
    //             let i, isum = 0
    //             for(i = 0; i < 1e10; i++){
    //                 isum += i
    //             }
    //             // console.log(isum)
    //             console.log("successQ!!")
    //             return isum
            
    //             }
    //         )
    //     }
    // })
    // return Promise.reject().then(() => {
    //     let i, isum = 0
    //     for(i = 0; i < 1e9; i++){
    //         isum += i
    //     }
    //     // console.log(isum)
    //     return isum
    // })
}
// async 
// async 
async function getTowork(){
    // ()
    longProcess().then(e => console.log(e, 777)) 
    let lon =  longProcess() 
    let lin = longProcess()

    res = await Promise.all([lin, lon])
    console.log(res, "resolute")
    // longProcess().then(e => console.log(e()))
    
    console.log("Other Stuffs...")
    console.log("Other Stuffs...")

    console.log(lin, lon)

    console.log("Other Stuffs...")
    // return Promise.all([lin, lon])
    // let isum = longProcess().then(val => console.log(val())).catch((e) => console.error(e, "error expected!"))
    // console.log(isum)
    // for(let j = 0; j < 1e3; j++) console.log(j)

}

getTowork()
// getTowork().then((log) => console.log(log))
// longProcess(1
