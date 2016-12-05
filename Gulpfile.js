var gulp = require('gulp');
var concat = require('gulp-concat');
var cleanCSS = require('gulp-clean-css');
var uglify = require('gulp-uglify');
var rev = require('gulp-rev-hash');
var sourcemaps = require('gulp-sourcemaps');
var ts = require('gulp-typescript');

var bc = './bower_components/';

gulp.task('jslibs', function() {
    gulp.src([
        bc + 'angular/angular.min.js',
        bc + 'angular-sanitize/angular-sanitize.min.js',
        bc + 'ng-file-upload/ng-file-upload.min.js',
        bc + 'angular-bootstrap/ui-bootstrap-tpls.min.js',  // todo: pagination only
        bc + 'angular-aria/angular-aria.min.js',
        bc + 'angular-animate/angular-animate.min.js',
        bc + 'angular-material/angular-material.min.js'
    ])
    .pipe(concat('libs.js'))
    .pipe(uglify())
    //.pipe(gulp.dest('./static/js/'));
    .pipe(gulp.dest('dist'));
});

gulp.task('csslibs', function () {
    gulp.src([
        bc + 'bootstrap/dist/css/bootstrap.min.css',
        bc + 'angular-material/angular-material.min.css'
    ])
    .pipe(concat('libs.css'))
    .pipe(gulp.dest('./static/css/'));
});

gulp.task('rev', function () {
    gulp.src('static/index.html')
        .pipe(rev())
        .pipe(gulp.dest('dist'));
});

gulp.task('revts', function () {
    gulp.src('static/index.ts.html')
        .pipe(rev())
        .pipe(gulp.dest('dist'));
});

gulp.task('compile', ['revts'], function () {
    gulp.src('static/typescript/*.ts')
        .pipe(sourcemaps.init())
        .pipe(ts({
            noImplicitAny: false,
            // out: 'output.js',
            target: 'ES5',
            removeComments: true
        }))
        .pipe(uglify())
        .pipe(sourcemaps.write())
        .pipe(gulp.dest('dist'));
});

gulp.task('dist', ['revts', 'minifycss', 'jslibs', 'compile']/*, function () {
    gulp.src('dist/*.js')
        .pipe(concat('build.js'))
        .pipe(gulp.dest('.'));
}*/);

gulp.task('clientlibs', ['jslibs', 'csslibs', 'revts']);


//==============================================================================


/*gulp.task('minifyjs', ['jslibs'], function () {
    gulp.src('./static/js/!*.js')
        .pipe(sourcemaps.init())
        .pipe(concat('app.js'))
        .pipe(uglify())
        .pipe(sourcemaps.write('dist'))
        .pipe(gulp.dest('dist'));
});*/

gulp.task('minifycss', ['csslibs'], function () {
    gulp.src('./static/css/*.css')
        .pipe(concat('app.css'))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('dist'));
});

gulp.task('build', ['minifyjs', 'minifycss']);



gulp.task('watch', function() {
    gulp.watch('static/typescript/*.ts', ['compile']);
    gulp.watch('static/index.ts.html', ['revts']);
    // gulp.watch('static/css/*.css', ['minifycss']);
});
