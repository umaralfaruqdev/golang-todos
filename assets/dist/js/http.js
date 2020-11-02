/**
 * Author: Umar
 * Created: 2020-10-26 17:44
 * Http request
 */

/**
 * handle click submit button
 * @return void
 */ 
 let handleSubmit = function (elem) {
    let submitType = elem.dataset.submitType;

    if (submitType === 'edit') {

    } else if (submitType === 'add') {
        sendData();
    } else {
        console.error("no have submit type");
    }
}

/**
 * post request to server
 * @return void
 */ 

let sendData = async function () {
    
    // get value
    let fname = _FORM.fname.value;
    let lname = _FORM.lname.value;
    let bornYear = _FORM.year.value;
    let bornMonth = _FORM.month.value;
    let bornDay = _FORM.day.value;
    let isMarried = _FORM.married.value;

    let bornString = `${bornYear}-${bornMonth}-${bornDay}`

    try {

        const req = await fetch('http://localhost:9000/user/store', {
            method: 'POST',
            headers: {
                'Content-Type': 'x-www-form-urlencoded',
            },
            body: JSON.stringify({
                fname: fname,
                lname: lname,
                born: bornString,
                is_married: isMarried,
            }),
        });

        // if successfuly statuc code
        let data = await req.json();
        await console.log(data);

    } catch (err) {
        // if catch
        console.error(err);
    }
}

/**
 * delete request to server
 * @return void
 */

const deleteUser = async function (id) {

    try {

        let data = {id: 12}

        const req = await fetch(`http://localhost:9000/user/delete?id=${id}`,)
        const res = await req.json();
        console.log(res)
    } catch (err) {
        console.error(err);
    }
}
