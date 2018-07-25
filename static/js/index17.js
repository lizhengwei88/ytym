;
(function() {
    $('.tab').on('click', 'a', function() {
        $(this).addClass('tb').siblings().removeClass('tb');
        $('section').find('main').stop().eq($(this).index()).animate({
            'left': '0%'
        }, 500).siblings().stop().animate({
            'left': '100%'
        }, 300)
    })
})()