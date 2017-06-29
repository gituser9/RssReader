var gulp = require('gulp');
var concat = require('gulp-concat');
var cleanCSS = require('gulp-clean-css');
var uglify = require('gulp-uglify');
var rev = require('gulp-rev-simple-hash');
var sourcemaps = require('gulp-sourcemaps');


var bc = './bower_components/';

gulp.task('jslibs', function() {
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
    ])
    .pipe(concat('libs.js'))
    .pipe(uglify())
    .pipe(gulp.dest('dist'));
});

gulp.task('csslibs', function() {
    gulp.src([
        bc + 'angular-material/angular-material.min.css'
    ])
    .pipe(concat('libs.css'))
    .pipe(gulp.dest('dist'));
});

gulp.task('revts', function() {
    gulp.src('static/html/index.html')
        .pipe(rev())
        .pipe(gulp.dest('dist'));
});

gulp.task('compile', function() {
    gulp.src('static/typescript/**/*.js')
        .pipe(sourcemaps.init())
        .pipe(concat('output.js'))
        .pipe(uglify())
        .pipe(sourcemaps.write('.'))
        .pipe(gulp.dest('dist'));
});

gulp.task('minifycss', function() {
    return gulp.src('static/css/*.css')
        .pipe(concat('app.css'))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('dist'));
});

gulp.task('dist', ['minifycss', 'jslibs', 'compile', 'revts']);

gulp.task('watch', function() {
    gulp.watch('static/typescript/*.js', ['compile']);
    gulp.watch('static/typescript/rss/*.js', ['compile']);
    gulp.watch('static/typescript/vk/*.js', ['compile']);
    gulp.watch('static/typescript/twitter/*.js', ['compile']);
    gulp.watch('static/typescript/models/*.js', ['compile']);
    gulp.watch('static/html/index.html', ['revts']);
    gulp.watch('static/css/*.css', ['minifycss']);
});
