const gulp = require('gulp');
const concat = require('gulp-concat');
const cleanCSS = require('gulp-clean-css');
const rev = require('gulp-rev-simple-hash');
const babel = require('gulp-babel');
const uglify      = require('gulp-uglify');
const sourcemaps  = require('gulp-sourcemaps');


const bc = './bower_components/';

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
gulp.task('dist', gulp.series( 'compile','minifycss', 'csslibs', 'jslibs', 'revts', function (done) {
    done();
}));
gulp.task('watch', function(){
    gulp.watch('static/js/**/*.js', gulp.series('jslibs', 'compile'))
    gulp.watch('static/html/index.html', gulp.series('revts'))
    gulp.watch('static/css/*.css', gulp.series('minifycss'))
    return;
});