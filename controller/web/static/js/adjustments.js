$(function(){
  var pageHeight;

  function set_columns() {
    pageHeight = 0;
    $('footer').css('margin-top', '0');
    $('.nav-border').css('height', '600');
    pageHeight = $(document).height();
    // console.log(pageHeight);
    $('.nav-border').css('height', pageHeight);

    var margin = pageHeight - 511 - 158;
    if (pageHeight < 800){
      margin = 270;
      $('body').css('height', '930');
      $('.nav-border').css('height', '930');
    }
    $('footer').css('margin-top', margin);
    $('footer').css('display', 'block');

    if($(window).width() > 1171){$('.docs-sidebar').css({'position': 'absolute', 'right': '0'})};
    if($(window).width() < 1171){$('.docs-sidebar').css('position', 'static');}
  }

  $(window).resize(set_columns);
});
