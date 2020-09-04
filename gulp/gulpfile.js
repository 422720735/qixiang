const path = require('path')
const gulp = require('gulp')
const sass = require('gulp-sass')
const cssUglify = require('gulp-minify-css')
const autoprefixer = require('gulp-autoprefixer')
const rename = require('gulp-rename')
const del = require('del')
gulp.task('sass', () => {
    return gulp.src('src/sass/*.scss')
        .pipe(sass())
        .pipe(autoprefixer({
            browsers: ['last 20 versions']
        }))
        .pipe(gulp.dest('src/css'))
})

gulp.task('cssUglify', function () {
    return gulp.src(['src/css/*.css'])
        .pipe(rename({suffix: '.min'}))
        .pipe(cssUglify())
        .pipe(gulp.dest('src/minCss'))
})

gulp.task('clear', function () {
    return del([
        path.resolve(__dirname, '../template/dist'),
    ], { force: true });
})

gulp.task('watch', () => {
    gulp.watch('src/sass/*.scss', gulp.series('sass'))
    gulp.watch('src/css/**', gulp.series('cssUglify'))
})

gulp.task('build', gulp.series('clear', 'cssUglify'))