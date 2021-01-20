jQuery(document).ready(function($) {
			 if ($(window).width() > 768) {	
	$('.header-menu-con li').hover(function() {
			  $(this).children('ul').show();
        },
        function() {
			$(this).children('ul').hide();
        });
	 }

 
	$('.proli span').click(function() {
			  $(this).parents(".proli").children('ul').slideToggle();
			  $(this).parents(".proli").find('h3').toggleClass('hover');
			  $(this).parents(".proli").siblings().children('ul').slideUp();
			  $(this).parents(".proli").siblings().find('h3').removeClass('hover');
        }),

	 
$('#slider .owl-carousel').owlCarousel({
    loop:true,
	items: 1,
	autoplay:true,
	autoplayTimeout:5000,
	autoplayHoverPause:true,//nav:true,navigationText:["prev","next"]
})
$('.youshi-con .owl-carousel').owlCarousel({
    loop:true,
	items: 1,
	autoplay:true,
	autoplayTimeout:7000,
	autoplayHoverPause:true,
})

$('.cp-img .owl-carousel').owlCarousel({
    loop:true,
	items: 1,
	autoplay:true,
	autoplayTimeout:5000,
	autoplayHoverPause:true,
})
$('.case-con .owl-carousel').owlCarousel({
    loop:true,
	items: 4,
	autoplay:true,
	autoplayTimeout:8000,
	autoplayHoverPause:true,
	margin:20,
	responsiveClass:true,
    responsive:{
        0:{
            items:1,
        },
        600:{
            items:2,

        },
        1000:{
            items:4,

            loop:false
        }
    }
})
	$('.entry-content img').parent("a").addClass("fancybox").attr("data-fancybox-group","gallery");
	$('.fancybox').fancybox();	
	$('#close_im').bind('click',function(){
		$('#main-im').css("height","0");
		$('#im_main').hide();
		$('#open_im').show();
	});
	$('#open_im').bind('click',function(e){
		$('#main-im').css("height","272");
		$('#im_main').show();
		$(this).hide();
	});
	$('.go-top').bind('click',function(){
		$(window).scrollTop(0);
	});

$('#header .button').on('click', function() {
			if ($(this).toggleClass('active').hasClass('active')) {
				$('.header-menu-con').addClass('active');
			} else {
				$('.header-menu-con').removeClass('active');
			}
		});
});