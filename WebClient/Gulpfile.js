var gulp = require('gulp');
var concat = require('gulp-concat');
var cleanCSS = require('gulp-clean-css');
var uglify = require('gulp-uglify');
var rev = require('gulp-rev-simple-hash');
var sourcemaps = require('gulp-sourcemaps');



var babel = require('gulp-babel');
// var browserify = require('gulp-browserify');
var browserify  = require('browserify');
var babelify    = require('babelify');
var source      = require('vinyl-source-stream');
var buffer      = require('vinyl-buffer');
var uglify      = require('gulp-uglify');
var sourcemaps  = require('gulp-sourcemaps');
var livereload  = require('gulp-livereload');


var bc = './bower_components/';

gulp.task('jslibs', gulp.series(function(done) {
    gulp.src([
        bc + 'jquery/dist/jquery.js',
        bc + 'angular/angular.min.js',
        bc + 'angular-sanitize/angular-sanitize.min.js',
        bc + 'ng-file-upload/ng-file-upload.min.js',
        bc + 'angular-aria/angular-aria.min.js',
        bc + 'angular-animate/angular-animate.min.js',
        bc + 'angular-material/angular-material.min.js',
        bc + 'ngInfiniteScroll/build/ng-infinite-scroll.min.js',
        bc + 'angular-paging/dist/paging.js'
    ], { allowEmpty: true })
    .pipe(concat('libs.js'))
    .pipe(uglify())
    .pipe(gulp.dest('dist'));
    done();
}));
gulp.task('csslibs', gulp.series(function(done) {
    gulp.src([
        bc + 'angular-material/angular-material.min.css'
    ])
    .pipe(concat('libs.css'))
    .pipe(gulp.dest('dist'));
    done();
}));
gulp.task('revts', gulp.series(function(done) {
    gulp.src('static/html/index.html')
        .pipe(rev())
        .pipe(gulp.dest('dist'));
    done();
}));
gulp.task('compile', gulp.series(function(done) {
    gulp.src('static/js/**/*.js')
        .pipe(sourcemaps.init())
        .pipe(babel({
            presets: ['@babel/env']
        }))
        .pipe(concat('output.js'))
        .pipe(uglify())
        .pipe(sourcemaps.write('.'))
        .pipe(gulp.dest('dist'));
    done();
}));
gulp.task('minifycss', gulp.series(function(done) {
    return gulp.src('static/css/*.css')
        .pipe(concat('app.css'))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('dist'));
    done();
}));
gulp.task('dist', gulp.series('minifycss', 'csslibs', 'jslibs', 'compile', 'revts', function (done) {
    done();
}));
gulp.task('watch', function(){
    gulp.watch('static/js/**/*.js', gulp.series('jslibs', 'compile')),
    gulp.watch('static/html/index.html', gulp.series('revts')),
    gulp.watch('static/css/*.css', gulp.series('minifycss'));
    return;
});