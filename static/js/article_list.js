function select_list(click_type, click_id) {
    let ids = document.getElementsByClassName('active');
    let type_id, provice_id, city_id, sort_id;
    for (const id of ids){
        if (id.title === 'type'){
            type_id = id.name;
        } else if (id.title === 'provice') {
            provice_id  = id.name;
        } else if (id.title === 'city') {
            city_id  = id.name;
        } else if (id.title === 'sort') {
            sort_id  = id.name;
        }
    }

    let url_str; //  /publish/t/provice/33/city/191
    let url_arr = [sort_id, type_id, 'provice', provice_id, 'city', city_id];
    if (city_id === undefined || city_id === '0') {
        url_arr.splice(4);
    }
    if (provice_id === '0') {
        url_arr.splice(2);
    }

    if (click_type === 'msg_type'){
        url_arr[1] = click_id;
    }else if (click_type === 'sort_type'){
        url_arr[0] = click_id;
    }else if (click_type === 'provice'){
        if (url_arr.length === 2){
            url_arr.push('provice', click_id)
        }else if (url_arr.length === 4){
            url_arr[3] = click_id;
        } else if (url_arr.length === 6){
            url_arr[3] = click_id;
            url_arr.splice(4);
        }

        if (click_id === '0'){
            url_arr.splice(2);
        }
    }else if (click_type === 'city'){
        if (url_arr.length === 4){
            url_arr.push('city', click_id);
        }else if (url_arr.length === 6){
            url_arr[5] = click_id;
        }

        if (click_id === '0') {
            url_arr.splice(4);
        }
    }
    url_str = "/" + url_arr.join("/");
    window.location.href = url_str;
}


function go_page(page_num) {
    let ids = document.getElementsByClassName('active');
    let type_id, provice_id, city_id, sort_id;
    for (const id of ids){
        if (id.title === 'type'){
            type_id = id.name;
        } else if (id.title === 'provice') {
            provice_id  = id.name;
        } else if (id.title === 'city') {
            city_id  = id.name;
        } else if (id.title === 'sort') {
            sort_id  = id.name;
        }
    }

    let url_str; //  /publish/t/provice/33/city/191
    let url_arr = [sort_id, type_id, 'provice', provice_id, 'city', city_id];

    if (city_id === undefined || city_id === '0') {
        url_arr.splice(4);
    }
    if (provice_id === '0') {
        url_arr.splice(2);
    }

    url_arr.push('page', page_num);
    url_str = "/" + url_arr.join("/");
    window.location.href = url_str;
}